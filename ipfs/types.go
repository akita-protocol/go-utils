package ipfs

type Arc3 struct {
	Name           string     `json:"name"`
	Decimals       int64      `json:"decimals"`
	Description    string     `json:"description"`
	Image          string     `json:"image"`
	ImageIntegrity string     `json:"image_integrity"`
	ImageMimetype  string     `json:"image_mimetype"`
	UnitName       string     `json:"unitName"`
	AssetName      string     `json:"assetName"`
	Properties     Properties `json:"properties"`
}

type Properties struct {
	Creator         Creator       `json:"creator"`
	Royalties       []Royalty     `json:"royalties"`
	Collection      Collection    `json:"collection"`
	KeyWords        []interface{} `json:"keyWords"`
	Publisher       string        `json:"publisher"`
	ItemListElement int64         `json:"itemListElement"`
	NumberOfItems   int64         `json:"numberOfItems"`
	Arc69           Arc69         `json:"arc69"`
}

type Arc69 struct {
	Standard   string        `json:"standard"`
	Attributes []interface{} `json:"attributes"`
}

type Collection struct {
	Name string `json:"name"`
}

type Creator struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
}

type Royalty struct {
	Name  string `json:"name"`
	Addr  string `json:"addr"`
	Share int64  `json:"share"`
}

type Arc19 struct {
	Name           string             `json:"name"`
	Description    *string            `json:"description"`
	Standard       string             `json:"standard"`
	Decimals       *int64             `json:"decimals"`
	Image          string             `json:"image"`
	ImageMimetype  string             `json:"image_mimetype"`
	ImageIntegrity *string            `json:"image_integrity"`
	Properties     *map[string]string `json:"properties"`
}
