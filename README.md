# SUZURI API v1 for golang

## Usage
```
$ git clone https://github.com/go-numb/go-suzuri.git

```

```
func main() {
    // let's get APIKEY at https://suzuri.jp/developer/apps
    // Clientを作る
	c := New(<api_key>)

    // 現在提供されているSUZURI内のProductItemsを取得する
	items, err := c.GetItems()
	if err != nil {
		log.Error(err)
	}

    // NewMaterial() で 制作するtitle, filename, 取得したitemsをmaterialにわたす
    // ここではすべて渡しているが、itemsから必要なものだけ渡すことも可能
	material := c.NewMaterial("testしています", "https://41.media.tumblr.com/QA9JpdgnOc8ov98s1C4S5EjJ_500.jpg", items)

	/* 必要性に応じて
	for i, item := range material.Products {
		material[i].ResizeMode = "contain" or "cover"
	}

	material.Description: "説明文",	
	*/

    // Createにmaterialを渡し、Uploadして商品化する
	if err := c.Create(material); err != nil {
		log.Error(err)
	}


	/*
		# For only text
	*/
	p, err := c.CreateByText("SUZURI 完全に理解した")
	if err != nil {
		log.Error(err)
	}

	fmt.Printf("%+v\n", p)
    
}
```