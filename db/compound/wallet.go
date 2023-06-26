package compound

import "gitlab.algotoolchain.com/keelakay/utils/db"

type Wallet struct {
	Address string  `json:"address"`
	NFD     *db.NFD `json:"nfd,omitempty"`
}
