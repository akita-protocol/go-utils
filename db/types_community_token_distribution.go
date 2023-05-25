package db

type CommunityTokenDistribution struct {
	AsaID     uint64 `structs:"asa_id,omitempty" db:"asa_id" json:"asa_id"`
	Frequency string `structs:"frequency,omitempty" db:"frequency" json:"frequency,omitempty"`
}
