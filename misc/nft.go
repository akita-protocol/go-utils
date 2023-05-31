package misc

type NFTPreview struct {
	AsaID        uint64 `json:"asa_id"`
	Creator      string `json:"creator"`
	Name         string `json:"name"`
	Mime         string `json:"mime"`
	UnitName     string `json:"unit_name"`
	URL          string `json:"url"`
	ImageURL     string `json:"image_url"`
	TotalCreated uint64 `json:"total_created"`
	Decimals     uint64 `json:"decimals"`
}
