package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type CollectionExcludedAssets struct {
	ID    string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	AsaID uint64 `structs:"asa_id,omitempty" db:"asa_id" json:"asa_id"`
}

func CollectionExcludedAssetsTableKeys() []string {
	return []string{"id", "asa_id"}
}

func GetCollectionExcludedAssets[H Handle](h H, id string) (*[]CollectionExcludedAssets, error) {
	const op errors.Op = "GetCollectionExcludedAssets"
	query := fmt.Sprintf("select %s from %s.collection_excluded_assets where id = ?", strings.Join(CollectionExcludedAssetsTableKeys(), ","), NFDDatabase())

	var assets []CollectionExcludedAssets
	err := h.Select(&assets, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Collection Excluded Assets Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &assets, nil
}

func GetCollectionExcludedAssetByAsaID[H Handle](h H, asaID uint64) (*[]CollectionExcludedAssets, error) {
	const op errors.Op = "GetCollectionExcludedAssetByAsaID"
	query := fmt.Sprintf("select %s from %s.collection_excluded_assets where asa_id = ?", strings.Join(CollectionExcludedAssetsTableKeys(), ","), NFDDatabase())

	var assets []CollectionExcludedAssets
	err := h.Select(&assets, query, asaID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Collection Excluded Assets Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &assets, nil
}
