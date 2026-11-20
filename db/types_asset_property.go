package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/akita-protocol/go-utils/errors"
	"github.com/akita-protocol/go-utils/misc"
)

type AssetProperty struct {
	ID        uint64 `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	MetaKey   string `structs:"meta_key,omitempty" db:"meta_key" json:"meta_key,omitempty"`
	MetaValue string `structs:"meta_value,omitempty" db:"meta_value" json:"meta_value,omitempty"`
}

func AssetPropertyTableKeys() []string {
	return []string{"id", "meta_key", "meta_value"}
}

func GetAssetProperty[H Handle](h H, id uint64, metaKey string) (*AssetProperty, error) {
	const op errors.Op = "GetAssetProperty"
	query := fmt.Sprintf("select %s from %s.asset_properties where id = ? and meta_key = ?", strings.Join(AssetPropertyTableKeys(), ","), NFDDatabase())

	var a AssetProperty
	err := h.Get(&a, query, id, metaKey)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "AssetProperty Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &a, nil
}

func GetAssetProperties[H Handle](h H, id uint64) (*[]AssetProperty, error) {
	const op errors.Op = "GetAssetProperties"
	query := fmt.Sprintf("select %s from %s.asset_properties where id = ?", strings.Join(AssetPropertyTableKeys(), ","), NFDDatabase())

	var a []AssetProperty
	err := h.Select(&a, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "AssetProperty Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &a, nil
}

func GetAssetPropertiesIn[H Handle](h H, ids []uint64) (*[]AssetProperty, error) {
	const op errors.Op = "GetAssetPropertiesIn"
	query := fmt.Sprintf("select %s from %s.asset_properties where id in (%s)", strings.Join(AssetPropertyTableKeys(), ","), NFDDatabase(), strings.Repeat("?, ", len(ids))[0:(len(ids)*3)-2])

	var a []AssetProperty
	err := h.Select(&a, query, misc.ToInterfaceSlice(ids)...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "AssetProperty Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &a, nil
}

func GetAssetPropertiesInAsMap[H Handle](h H, ids []uint64) (map[string][]string, error) {
	const op errors.Op = "GetAssetPropertiesInAsMap"

	a, err := GetAssetPropertiesIn(h, ids)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}

	m := make(map[string][]string)
	for i := range *a {
		v := (*a)[i]
		m[v.MetaKey] = append(m[v.MetaKey], v.MetaValue)
	}

	return m, nil
}

func DeleteAssetPropertiesNotIn[H Handle](h H, id uint64, metaKeys []string) error {
	const op errors.Op = "DeleteAssetPropertiesNotIn"
	query := fmt.Sprintf("delete from %s.asset_properties where id = ? and meta_key not in (%s)", NFDDatabase(), strings.Repeat("?, ", len(metaKeys))[0:(len(metaKeys)*3)-2])
	_, err := h.Exec(query, append([]interface{}{id}, misc.ToInterfaceSlice(metaKeys)...)...)
	if err != nil {
		return errors.E(pkg, op, err)
	}
	return nil
}
