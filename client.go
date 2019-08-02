package suzuri

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/gommon/log"
)

const (
	URI     = "https://suzuri.jp/api"
	VERSION = "v1"
)

type Client struct {
	Key        string
	HTTPClient *http.Client

	// 200/15minutes 2019/08/01現在
	API *API
}

type API struct {
	sync.Mutex

	Limit  int
	Remain int
	Reset  time.Time
}

func New(apiKey string) *Client {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &Client{
		Key:        apiKey,
		HTTPClient: client,
		API: &API{
			Limit:  200,
			Remain: 200,
			Reset:  time.Now(),
		},
	}
}

func (p *API) isLimit() bool {
	p.Lock()
	defer p.Unlock()

	if p.Remain == 0 {
		return true
	}

	return false
}

func (p *API) get(h http.Header) {
	p.Lock()
	defer p.Unlock()

	limit, err := strconv.Atoi(h.Get("X-RateLimit-Limit"))
	if err != nil {
		return
	}
	p.Limit = limit
	remain, err := strconv.Atoi(h.Get("X-RateLimit-Remaining"))
	if err != nil {
		return
	}
	p.Remain = remain

	// TODO: string time to time.Time 格納
	log.Debugf("reset time: %+v\n", h.Get("X-RateLimit-Reset"))
	// layout := ""
	// p.Remain, _ = time.Parse(h.Get("X-RateLimit-Reset"))
}

func (p *API) toString() string {
	return fmt.Sprintf("limit: %d, remain: %d, reset time: %v", p.Limit, p.Remain, p.Reset)
}

func (p *Client) request(method, endpoint string, body []byte) (*http.Request, error) {
	if p.API.isLimit() {
		return nil, fmt.Errorf("api limit, is %d", p.API.Remain)
	}

	u, err := url.ParseRequestURI(fmt.Sprintf("%s/%s/%s", URI, VERSION, endpoint))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		method,
		u.String(),
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}

	// sets Authorization
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.Key))
	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (p *Client) do(req *http.Request, in interface{}) error {
	res, err := p.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	p.API.get(res.Header)
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%s, %s",
			res.Status,
			p.API.toString())
	}

	if err := json.NewDecoder(res.Body).Decode(in); err != nil {
		return err
	}

	return nil
}
