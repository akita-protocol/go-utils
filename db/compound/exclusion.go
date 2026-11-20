package compound

import "github.com/akita-protocol/go-utils/db"

type StakingExclusion struct {
	ID    string            `json:"id,omitempty"`
	Pools *[]db.StakingPool `json:"pools,omitempty"`
}
