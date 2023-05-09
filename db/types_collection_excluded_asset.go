package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type CollectionExcludedAsset struct {
	ID    string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	AsaID uint64 `structs:"asa_id,omitempty" db:"asa_id" json:"asa_id"`
}

func CollectionExcludedAssetTableKeys() []string {
	return []string{"id", "asa_id"}
}

func GetCollectionExcludedAssets[H Handle](h H, id string) (*[]CollectionExcludedAsset, error) {
	const op errors.Op = "GetCollectionExcludedAssets"
	query := fmt.Sprintf("select %s from %s.collection_excluded_assets where id = ?", strings.Join(CollectionExcludedAssetTableKeys(), ","), NFDDatabase())

	var assets []CollectionExcludedAsset
	err := h.Select(&assets, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Collection Excluded Assets Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &assets, nil
}

func GetCollectionExcludedAssetByAsaID[H Handle](h H, asaID uint64) (*[]CollectionExcludedAsset, error) {
	const op errors.Op = "GetCollectionExcludedAssetByAsaID"
	query := fmt.Sprintf("select %s from %s.collection_excluded_assets where asa_id = ?", strings.Join(CollectionExcludedAssetTableKeys(), ","), NFDDatabase())

	var assets []CollectionExcludedAsset
	err := h.Select(&assets, query, asaID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Collection Excluded Assets Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &assets, nil
}

// DeletCollectionExcludedAssetsNotIn deletes collection excluded assets that are not included in a list for a given collection
func DeletCollectionExcludedAssetsNotIn[H Handle](h H, data ...interface{}) error {
	const op errors.Op = "DeletCollectionExcludedAssetsNotIn"
	qMarks := []rune(strings.Repeat("?, ", (len(data) - 1)))
	query := "delete from %s.collection_excluded_assets where id = ? and id not in (" + string(qMarks[0:len(qMarks)-2]) + ")"
	var err error

	switch h := any(h).(type) {
	case *sqlx.Tx:
		stmt, err := h.Prepare(query)
		if err != nil {
			return errors.E(pkg, op, errors.Database, err, "Failed to Prepare Query")
		}

		_, err = stmt.Exec(data...)
		if err != nil {
			return errors.E(pkg, op, errors.Database, err, "Failed to Execute Query")
		}
	case *sqlx.DB:
		_, err = h.Exec(query, data...)
		if err != nil {
			return errors.E(pkg, op, errors.Database, err, "Failed to Execute Query")
		}
	}

	return nil
}
