package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type AssetGroup struct {
	ID      string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	AssetID uint64 `structs:"asset_id,omitempty" db:"asset_id" json:"asset_id,omitempty"`
	LP      *bool  `structs:"lp,omitempty" db:"lp" json:"lp,omitempty"`
}

func AssetGroupTableKeys() []string {
	return []string{"id", "asset_id", "lp"}
}

func GetAssetGroup[H DBStruct](h H, id string) (*[]AssetGroup, error) {
	const op errors.Op = "GetAssetGroup"
	query := fmt.Sprintf("select %v from %s.asset_group where id = ?", strings.Join(AssetGroupTableKeys(), ", "), DiscordBotDatabase())
	var list []AssetGroup

	err := h.Select(&list, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "AssetGroup Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	if !(len(list) > 0) {
		return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, fmt.Errorf("AssetGroup Not Found"))
	}

	return &list, nil
}
