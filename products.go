package suzuri

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type ResponseFromProducts struct {
	Products []struct {
		ID             int       `json:"id"`
		Title          string    `json:"title"`
		Published      bool      `json:"published"`
		PublishedAt    time.Time `json:"publishedAt"`
		CreatedAt      time.Time `json:"createdAt"`
		UpdatedAt      time.Time `json:"updatedAt"`
		ExamplaryAngle string    `json:"examplaryAngle"`
		ImageURL       string    `json:"imageUrl"`
		SampleImageURL string    `json:"sampleImageUrl"`
		URL            string    `json:"url"`
		SampleURL      string    `json:"sampleUrl"`
		Item           struct {
			ID           int           `json:"id"`
			Name         string        `json:"name"`
			Angles       []interface{} `json:"angles"`
			HumanizeName string        `json:"humanizeName"`
		} `json:"item"`
		Material struct {
			ID             int       `json:"id"`
			Title          string    `json:"title"`
			Description    string    `json:"description"`
			Price          int       `json:"price"`
			Violation      bool      `json:"violation"`
			Published      bool      `json:"published"`
			PublishedAt    time.Time `json:"publishedAt"`
			UploadedAt     time.Time `json:"uploadedAt"`
			DominantRgb    string    `json:"dominantRgb"`
			OriginalWidth  int       `json:"originalWidth"`
			OriginalHeight int       `json:"originalHeight"`
			User           struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				DisplayName string `json:"displayName"`
				AvatarURL   string `json:"avatarUrl"`
			} `json:"user"`
		} `json:"material"`
		SampleItemVariant struct {
			ID        int  `json:"id"`
			Price     int  `json:"price"`
			Exemplary bool `json:"exemplary"`
			Color     struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Rgb  string `json:"rgb"`
			} `json:"color"`
			Size struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"size"`
		} `json:"sampleItemVariant"`
	} `json:"products"`
	Meta struct {
		HasNext bool `json:"hasNext"`
	} `json:"meta"`
}

// GetProducts 自他商品群取得
func (p *Client) GetProducts(userID, itemID, materialID, userName string) (*ResponseFromProducts, error) {
	v := url.Values{}
	if userID != "" {
		v.Set("userId", userID)
	}
	if itemID != "" {
		v.Set("itemId", itemID)
	}
	if materialID != "" {
		v.Set("materialId", materialID)
	}
	if userName != "" {
		v.Set("userName", userName)
	}

	q := v.Encode()
	if q != "" {
		q = "?" + q
	}

	req, err := p.request(
		http.MethodGet,
		"products"+q,
		nil,
	)
	if err != nil {
		return nil, err
	}

	var s = new(ResponseFromProducts)
	if err := p.do(req, s); err != nil {
		return nil, err
	}

	return s, nil
}

type ParamsForCreate struct {
	// Required
	Texture string `json:"texture,omitempty"`
	Title   string `json:"title,omitempty"`

	// Options
	Price       int       `json:"price,omitempty"`
	Description string    `json:"description,omitempty"`
	Products    []Product `json:"products,omitempty"`
}

type Product struct {
	ItemID                 int           `json:"itemId,omitempty"`
	ExemplaryItemVariantID int           `json:"exemplaryItemVariantId,omitempty"`
	Published              bool          `json:"published,omitempty"`
	ResizeMode             string        `json:"resizeMode,omitempty"`
	SubMaterials           []SubMaterial `json:"sub_materials,omitempty"`
}

type SubMaterial struct {
	Texture   string `json:"texture,omitempty"`
	PrintSide string `json:"printSide,omitempty"`
	Enabled   bool   `json:"enabled,omitempty"`
}

type ResponseFromCreate struct {
	Material struct {
		ID             int       `json:"id,omitempty"`
		Title          string    `json:"title,omitempty"`
		Description    string    `json:"description,omitempty"`
		Price          int       `json:"price,omitempty"`
		Violation      bool      `json:"violation,omitempty"`
		Published      bool      `json:"published,omitempty"`
		PublishedAt    time.Time `json:"publishedAt,omitempty"`
		UploadedAt     time.Time `json:"uploadedAt,omitempty"`
		DominantRgb    string    `json:"dominantRgb,omitempty"`
		OriginalWidth  int       `json:"originalWidth,omitempty"`
		OriginalHeight int       `json:"originalHeight,omitempty"`
		User           struct {
			ID          int    `json:"id,omitempty"`
			Name        string `json:"name,omitempty"`
			DisplayName string `json:"displayName,omitempty"`
			AvatarURL   string `json:"avatarUrl,omitempty"`
		} `json:"user,omitempty"`
	} `json:"material,omitempty"`
	Products []struct {
		ID             int       `json:"id,omitempty"`
		Title          string    `json:"title,omitempty"`
		Published      bool      `json:"published,omitempty"`
		PublishedAt    time.Time `json:"publishedAt,omitempty"`
		CreatedAt      time.Time `json:"createdAt,omitempty"`
		UpdatedAt      time.Time `json:"updatedAt,omitempty"`
		ExamplaryAngle string    `json:"examplaryAngle,omitempty"`
		ImageURL       string    `json:"imageUrl,omitempty"`
		SampleImageURL string    `json:"sampleImageUrl,omitempty"`
		URL            string    `json:"url,omitempty"`
		SampleURL      string    `json:"sampleUrl,omitempty"`
		Item           struct {
			ID           int           `json:"id,omitempty"`
			Name         string        `json:"name,omitempty"`
			Angles       []interface{} `json:"angles,omitempty"`
			HumanizeName string        `json:"humanizeName,omitempty"`
		} `json:"item,omitempty"`
		Material struct {
			ID             int       `json:"id,omitempty"`
			Title          string    `json:"title,omitempty"`
			Description    string    `json:"description,omitempty"`
			Price          int       `json:"price,omitempty"`
			Violation      bool      `json:"violation,omitempty"`
			Published      bool      `json:"published,omitempty"`
			PublishedAt    time.Time `json:"publishedAt,omitempty"`
			UploadedAt     time.Time `json:"uploadedAt,omitempty"`
			DominantRgb    string    `json:"dominantRgb,omitempty"`
			OriginalWidth  int       `json:"originalWidth,omitempty"`
			OriginalHeight int       `json:"originalHeight,omitempty"`
			User           struct {
				ID          int    `json:"id,omitempty"`
				Name        string `json:"name,omitempty"`
				DisplayName string `json:"displayName,omitempty"`
				AvatarURL   string `json:"avatarUrl,omitempty"`
			} `json:"user,omitempty"`
		} `json:"material,omitempty"`
		SampleItemVariant struct {
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
		} `json:"sampleItemVariant,omitempty"`
	} `json:"products,omitempty"`
}

// NewMaterial SUZURI提供商品群選択
// 現行は全選択
func (p *Client) NewMaterial(title, filename string, toribun int, items []Item) *ParamsForCreate {
	l := len(items)
	products := make([]Product, l)
	for i, item := range items {
		products[i].ItemID = item.ID
		products[i].Published = true
		// products[i].ResizeMode = "contain"
		// products[i].ExemplaryItemVariantID = 151
		// products[i].SubMaterials = []SubMaterial{SubMaterial{Enabled: false}}
	}

	product := &ParamsForCreate{
		Title:    title,
		Texture:  filename,
		Price:    toribun,
		Products: products,
	}

	return product
}

// Create 指定画像で商品群を登録
func (p *Client) Create(params *ParamsForCreate) (*ResponseFromCreate, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := p.request(
		http.MethodPost,
		"materials",
		body,
	)
	if err != nil {
		return nil, err
	}

	var s = new(ResponseFromCreate)
	if err := p.do(req, s); err != nil {
		return nil, err
	}

	return s, nil
}
