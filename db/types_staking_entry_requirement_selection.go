package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type StakingEntryRequirementSelection struct {
	EntryRequirementID string `structs:"er_id,omitempty" db:"er_id" json:"entry_requirement_id,omitempty"`
	// GroupID            *string `structs:"group_id,omitempty" db:"group_id" json:"group_id,omitempty"`
	Quantity uint64 `structs:"quantity,omitempty" db:"quantity" json:"quantity,omitempty"`
}

func EntryRequirementSelectionTableKeys() []string {
	return []string{"er_id", "quantity"}
}

func GetStakingEntryRequirementSelection[H Handle](h H, erID string) (*StakingEntryRequirementSelection, error) {
	const op errors.Op = "GetStakingEntryRequirementSelection"
	query := fmt.Sprintf("select %s from %s.entry_requirement_selection where er_id = ?", strings.Join(EntryRequirementSelectionTableKeys(), ","), StakingDatabase())

	var p StakingEntryRequirementSelection
	err := h.Get(&p, query, erID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Pool Entry Requirement Selection Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &p, nil
}
