package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/akita-protocol/go-utils/errors"
	"github.com/akita-protocol/go-utils/misc"
)

type ShuffleAssetProperty struct {
	ID        uint64 `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	MetaKey   string `structs:"meta_key,omitempty" db:"meta_key" json:"meta_key,omitempty"`
	MetaValue string `structs:"meta_value,omitempty" db:"meta_value" json:"meta_value,omitempty"`
}

func ShuffleAssetPropertyTableKeys() []string {
	return []string{"id", "meta_key", "meta_value"}
}

func GetShuffleAssetProperty[H Handle](h H, id uint64, metaKey string) (*ShuffleAssetProperty, error) {
	const op errors.Op = "GetAssetProperty"
	query := fmt.Sprintf("select %s from %s.shuffle_asset_property where id = ? and meta_key = ?", strings.Join(AssetPropertyTableKeys(), ","), ShuffleDatabase())

	var a ShuffleAssetProperty
	err := h.Get(&a, query, id, metaKey)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "ShuffleAssetProperty Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &a, nil
}

func GetShuffleAssetProperties[H Handle](h H, id uint64) (*[]ShuffleAssetProperty, error) {
	const op errors.Op = "GetShuffleAssetProperties"
	query := fmt.Sprintf("select %s from %s.shuffle_asset_property where id = ?", strings.Join(AssetPropertyTableKeys(), ","), ShuffleDatabase())

	var a []ShuffleAssetProperty
	err := h.Select(&a, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "AssetProperty Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &a, nil
}

func DeleteShuffleAssetPropertiesNotIn[H Handle](h H, id uint64, metaKeys []string) error {
	const op errors.Op = "DeleteAssetPropertiesNotIn"
	query := fmt.Sprintf("delete from %s.shuffle_asset_property where id = ? and meta_key not in (%s)", ShuffleDatabase(), strings.Repeat("?, ", len(metaKeys))[0:(len(metaKeys)*3)-2])
	_, err := h.Exec(query, append([]interface{}{id}, misc.ToInterfaceSlice(metaKeys)...)...)
	if err != nil {
		return errors.E(pkg, op, err)
	}
	return nil
}
