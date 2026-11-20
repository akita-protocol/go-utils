package db

import (
	"database/sql"
	"fmt"

	"github.com/akita-protocol/go-utils/errors"
)

type SubscriptionFactory struct {
	ID               uint64 `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Version          uint64 `structs:"version,omitempty" db:"version" json:"version,omitempty"`
	CurrentAsOfBlock uint64 `structs:"current_as_of_block,omitempty" db:"current_as_of_block" json:"current_as_of_block,omitempty"`
}

func GetSubscriptionFactoryByID[H Handle](h H, id uint64) (*SubscriptionFactory, error) {
	const op errors.Op = "GetSubscriptionFactory"
	query := fmt.Sprintf("select id, version, current_as_of_block from %s.subscription_factory where id = ?", SubscriptionDatabase())
	var sf SubscriptionFactory

	err := h.Get(&sf, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "SubscriptionFactory Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &sf, nil
}

func GetLeastFactoryUpdateBlock[H Handle](h H) (uint64, error) {
	const op errors.Op = "GetLatestUpdateBlock"
	query := fmt.Sprintf("select min(current_as_of_block) from %s.subscription_factory", SubscriptionDatabase())
	var block uint64

	err := h.Get(&block, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "SubscriptionFactory Not Found")
		}
		return 0, errors.E(pkg, op, err)
	}
	return block, nil
}
