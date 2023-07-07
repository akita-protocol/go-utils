package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

// assets available in the shuffle
type ShuffleAsset struct {
	ShuffleID string `structs:"shuffle_id,omitempty" db:"shuffle_id" json:"shuffle_id,omitempty"`
	AssetID   uint64 `structs:"asset_id,omitempty" db:"asset_id" json:"asset_id,omitempty"`
}

func ShuffleAssetTableKeys() []string {
	return []string{
		"shuffle_id",
		"asset_id",
	}
}

func GetShuffleAssets[H Handle](h H, shuffleID string) (*[]ShuffleAsset, error) {
	const op errors.Op = "GetShuffleAssets"
	query := fmt.Sprintf("select %v from %v.shuffle_asset where shuffle_id = ?", strings.Join(ShuffleAssetTableKeys(), ","), ShuffleDatabase())

	var assets []ShuffleAsset
	err := h.Select(&assets, query, shuffleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Shuffle Asset Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &assets, err
}
