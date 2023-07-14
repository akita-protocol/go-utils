package db

import (
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type StakingPoolAsset struct {
	PoolID  string  `structs:"pool_id,omitempty" db:"pool_id" json:"pool_id,omitempty"`
	GroupID *string `structs:"group_id,omitempty" db:"group_id" json:"group_id,omitempty"`
	AssetID uint64  `structs:"asset_id,omitempty" db:"asset_id" json:"asset_id,omitempty"`
}

func StakingPoolAssetTableKeys() []string {
	return []string{"pool_id", "group_id", "asset_id"}
}

func GetStakingPoolAssetIDs[H Handle](h H, poolID string) (*[]uint64, error) {
	const op errors.Op = "GetPoolAssetIDs"
	query := fmt.Sprintf("select asset_id from %s.pool_asset where pool_id = ?", StakingDatabase())

	var assetIDs []uint64
	err := h.Select(&assetIDs, query, poolID)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}
	return &assetIDs, nil
}

func GetStakingPoolAssets[H Handle](h H, poolID string, st TableKeySlice) (*[]uint64, error) {
	const op errors.Op = "GetPoolAssetIDs"
	query := fmt.Sprintf("select %v from %v.asset where id in(select asset_id from %s.pool_asset where pool_id = ?)", strings.Join(AssetTableKeys(st), ","), NFDDatabase(), StakingDatabase())

	var assetIDs []uint64
	err := h.Select(&assetIDs, query, poolID)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}
	return &assetIDs, nil
}
