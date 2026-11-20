package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/akita-protocol/go-utils/errors"
)

type StakingEntryRequirementType string

const (
	StakingEntryRequirementTypeSelection StakingEntryRequirementType = "selection"
	StakingEntryRequirementTypeAsset     StakingEntryRequirementType = "asset"
)

type StakingEntryRequirement struct {
	ID     string                      `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	PoolID string                      `structs:"pool_id,omitempty" db:"pool_id" json:"pool_id,omitempty"`
	Type   StakingEntryRequirementType `structs:"type,omitempty" db:"type" json:"type,omitempty"`
}

func StakingEntryRequirementTableKeys() []string {
	return []string{"id", "pool_id", "type"}
}

func GetStakingEntryRequirement[H Handle](h H, id string) (*StakingEntryRequirement, error) {
	const op errors.Op = "GetStakingEntryRequirement"
	query := fmt.Sprintf("select %s from %s.entry_requirement where id = ?", strings.Join(StakingEntryRequirementTableKeys(), ","), StakingDatabase())

	var p StakingEntryRequirement
	err := h.Get(&p, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Pool Entry Requirements Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &p, nil
}

func GetStakingEntryRequirementsByPoolID[H Handle](h H, poolID string) (*[]StakingEntryRequirement, error) {
	const op errors.Op = "GetStakingEntryRequirementsByPoolID"
	query := fmt.Sprintf("select %s from %s.entry_requirement where pool_id = ?", strings.Join(StakingEntryRequirementTableKeys(), ","), StakingDatabase())

	var list []StakingEntryRequirement
	err := h.Select(&list, query, poolID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Pool Entry Requirements Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}
