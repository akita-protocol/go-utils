package db

import (
	"database/sql"
	"fmt"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type StakingExclusion struct {
	ID     string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	PoolID string `structs:"pool_id,omitempty" db:"pool_id" json:"pool_id,omitempty"`
}

func StakingExclusionTableKeys() []string {
	return []string{
		"id",
		"pool_id",
	}
}

func GetStakingExclusionPoolIDsByID[H Handle](h H, id string, excludeID string) (*[]string, error) {
	const op errors.Op = "GetStakingExclusionPoolIDsByID"
	query := fmt.Sprintf("select pool_id from %s.exclusion where id = ? and pool_id != ?", StakingDatabase())
	var poolIDs []string

	err := h.Select(&poolIDs, query, id, excludeID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Staking Exclusion Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &poolIDs, nil
}

func GetStakingExclusionIDsByPoolID[H Handle](h H, poolID string) (*[]string, error) {
	const op errors.Op = "GetStakingExclusionIDsByPoolID"
	query := fmt.Sprintf("select id from %s.exclusion where pool_id = ?", StakingDatabase())
	var ids []string
	err := h.Select(&ids, query, poolID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Staking Exclusion Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &ids, nil
}
