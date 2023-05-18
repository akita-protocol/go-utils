package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type StakingEntryRequirementAsset struct {
	EntryRequirementID string `structs:"er_id,omitempty" db:"er_id" json:"entry_requirement_id,omitempty"`
	AssetID            uint64 `structs:"asset_id,omitempty" db:"asset_id" json:"asset_id,omitempty"`
	Quantity           uint64 `structs:"quantity,omitempty" db:"quantity" json:"quantity,omitempty"`
}

func StakingEntryRequirementAssetTableKeys() []string {
	return []string{"er_id", "asset_id", "quantity"}
}

func GetStakingEntryRequirementAsset[H Handle](h H, erID string) (*StakingEntryRequirementAsset, error) {
	const op errors.Op = "GetStakingEntryRequirementAsset"
	query := fmt.Sprintf("select %s from %s.entry_requirement_asset where er_id = ?", strings.Join(StakingEntryRequirementAssetTableKeys(), ","), StakingDatabase())

	var p StakingEntryRequirementAsset
	err := h.Get(&p, query, erID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Pool Entry Requirement Asset Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &p, nil
}
