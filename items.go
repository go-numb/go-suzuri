package suzuri

import (
	"net/http"
)

type Items struct {
	Is []Item `json:"items"`
}

type Item struct {
	ID           int           `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Angles       []interface{} `json:"angles,omitempty"`
	HumanizeName string        `json:"humanizeName,omitempty"`
	Variants     []struct {
		ID        int  `json:"id,omitempty"`
		Price     int  `json:"price,omitempty"`
		Exemplary bool `json:"exemplary,omitempty"`
		Color     struct {
			ID   int    `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
			Rgb  string `json:"rgb,omitempty"`
		} `json:"color,omitempty"`
		Size struct {
			ID   int    `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"size,omitempty"`
	} `json:"variants,omitempty"`
}

func (p *Client) GetItems() ([]Item, error) {
	req, err := p.request(http.MethodGet, "items", nil)
	if err != nil {
		return nil, err
	}

	var items = new(Items)
	if err := p.do(req, items); err != nil {
		return nil, err
	}

	return items.Is, nil
}
