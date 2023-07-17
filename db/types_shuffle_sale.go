package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type ShuffleSale struct {
	Txn       string `structs:"txn,omitempty" db:"txn" json:"txn,omitempty"`
	ShuffleID string `structs:"shuffle_id,omitempty" db:"shuffle_id" json:"shuffle_id,omitempty"`
	AssetID   uint64 `structs:"asset_id,omitempty" db:"asset_id" json:"asset_id,omitempty"`
}

func ShuffleSaleTableKeys() []string {
	return []string{
		"txn",
		"shuffle_id",
		"asset_id",
	}
}

func GetShuffleSales[H Handle](h H, shuffleID string) (*[]ShuffleSale, error) {
	const op errors.Op = "GetShuffleSales"
	query := fmt.Sprintf("select %v from %v.shuffle_sale where shuffle_id = ?", strings.Join(ShuffleSaleTableKeys(), ","), ShuffleDatabase())

	var sales []ShuffleSale
	err := h.Select(&sales, query, shuffleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Shuffle Sale Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &sales, err
}

func GetShuffleSalesCount[H Handle](h H, shuffleID string) (*uint64, error) {
	const op errors.Op = "GetShuffleSales"
	query := fmt.Sprintf("select count(txn) from %v.shuffle_sale where shuffle_id = ?", ShuffleDatabase())

	var count uint64
	err := h.Get(&count, query, shuffleID)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}

	return &count, err
}

func GetShuffleSaleExists[H Handle](h H, shuffleID string, assetID uint64) (*bool, error) {
	const op errors.Op = "GetShuffleSaleExists"
	query := fmt.Sprintf("select exists(select 1 from %v.shuffle_sale where shuffle_id = ? and asset_id = ?)", ShuffleDatabase())

	var exists bool
	err := h.Get(&exists, query, shuffleID, assetID)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}

	return &exists, err
}
