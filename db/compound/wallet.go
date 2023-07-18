package compound

type Wallet struct {
	Address string `json:"address"`
	NFD     *NFD   `json:"nfd,omitempty"`
}
