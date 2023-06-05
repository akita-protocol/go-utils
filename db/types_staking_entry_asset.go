package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
)

type StakingEntryAsset struct {
	EntryID string `structs:"entry_id,omitempty" db:"entry_id" json:"entry_id,omitempty"`
	AssetID uint64 `structs:"asset_id,omitempty" db:"asset_id" json:"asset_id,omitempty"`
}

func StakingEntryAssetTableKeys() []string {
	return []string{"entry_id", "asset_id"}
}

func GetStakingEntryAssets[H Handle](h H, entryID string) (*[]uint64, error) {
	const op errors.Op = "GetEntryAssets"
	query := fmt.Sprintf("select asset_id from %s.entry_asset where entry_id = ?", StakingDatabase())
	var assets []uint64
	err := h.Select(&assets, query, entryID)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}
	return &assets, nil
}

func GetStakingEntryAssetsByUserIDAndPoolIDs[H Handle](h H, userID string, poolIDs []string) (*[]uint64, error) {
	const op errors.Op = "GetEntryAssetsByUserIDAndPoolIDs"
	query := fmt.Sprintf("select asset_id from %s.entry_asset where entry_id in (select id from %s.pool_entry where user_id = ? and txn is not null and pool_id in (%s))", StakingDatabase(), StakingDatabase(), strings.Repeat("?, ", len(poolIDs))[0:(len(poolIDs)*3)-2])
	var list []uint64

	args := []string{
		userID,
	}
	args = append(args, poolIDs...)

	err := h.Select(&list, query, misc.ToInterfaceSlice(args)...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Pool Entry Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &list, nil
}
