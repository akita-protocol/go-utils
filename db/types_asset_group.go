package db

import (
	"database/sql"

	"gitlab.keelakay.com/keelakay/utils/errors"
)

type AssetGroup struct {
	ID      string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	AssetID uint64 `structs:"asset_id,omitempty" db:"asset_id" json:"asset_id,omitempty"`
}

func GetAssetGroup[H DBStruct](h H, id string) (*[]uint64, error) {
	const op errors.Op = "GetAssetGroup"
	const query = "select asset_id from asset_group where id = ?"
	var list []uint64

	err := h.Select(&list, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "AssetGroup Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}
