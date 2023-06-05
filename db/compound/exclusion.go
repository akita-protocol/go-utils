package compound

import "gitlab.algotoolchain.com/keelakay/utils/db"

type StakingExclusion struct {
	ID    string            `json:"id,omitempty"`
	Pools *[]db.StakingPool `json:"pools,omitempty"`
}
