package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/akita-protocol/go-utils/errors"
)

// direct asset reservations for shuffle by address or whitelist
type ShuffleReserveAsset struct {
	ID string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	ShuffleID   string  `structs:"shuffle_id,omitempty" db:"shuffle_id" json:"shuffle_id,omitempty"`
	Address     *string `structs:"address,omitempty" db:"address" json:"address,omitempty"`
	WhitelistID *string `structs:"whitelist_id,omitempty" db:"whitelist_id" json:"whitelist_id,omitempty"`
	AssetID     uint64  `structs:"asset_id,omitempty" db:"asset_id" json:"asset_id,omitempty"`
}

func ShuffleReserveAssetTableKeys() []string {
	return []string{
		"id",
		"shuffle_id",
		"address",
		"whitelist_id",
		"asset_id",
	}
}

func GetShuffleReserveAssets[H Handle](h H, shuffleID string) (*[]ShuffleReserveAsset, error) {
	const op errors.Op = "GetShuffleReserveAssets"
	query := fmt.Sprintf("select %v from %v.shuffle_reserve_asset where shuffle_id = ?", strings.Join(ShuffleReserveAssetTableKeys(), ","), ShuffleDatabase())

	var assets []ShuffleReserveAsset
	err := h.Select(&assets, query, shuffleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Shuffle Reserve Asset Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &assets, err
}

func GetShuffleReserveAssetByWhitelistID[H Handle](h H, whitelistID string) (*ShuffleReserveAsset, error) {
	const op errors.Op = "GetShuffleReserveAssetByWhitelistID"
	query := fmt.Sprintf("select %v from %v.shuffle_reserve_asset where whitelist_id = ?", strings.Join(ShuffleReserveAssetTableKeys(), ","), ShuffleDatabase())

	var asset ShuffleReserveAsset
	err := h.Get(&asset, query, whitelistID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Shuffle Reserve Asset Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &asset, err
}
