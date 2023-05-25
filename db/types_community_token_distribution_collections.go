package db

type CommunityTokenDistributionCollection struct {
	AsaID  uint64 `structs:"asa_id,omitempty" db:"asa_id" json:"asa_id"`
	Name   string `structs:"name,omitempty" db:"name" json:"name,omitempty"`
	Amount uint64 `structs:"amount,omitempty" db:"amount" json:"amount,omitempty"`
}
