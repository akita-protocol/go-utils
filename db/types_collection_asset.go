package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type CollectionAsset struct {
	ID    string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	AsaID uint64 `structs:"asa_id,omitempty" db:"asa_id" json:"asa_id"`
}

func CollectionAssetTableKeys() []string {
	return []string{"id", "asa_id"}
}

func GetCollectionAssets[H Handle](h H, id string) (*[]CollectionAsset, error) {
	const op errors.Op = "GetCollectionAssets"
	query := fmt.Sprintf("select %s from %s.collection_asset where id = ?", strings.Join(CollectionAssetTableKeys(), ","), NFDDatabase())

	var assets []CollectionAsset
	err := h.Select(&assets, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Collection Assets Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &assets, nil
}

func GetCollectionAssetByAsaID[H Handle](h H, asaID uint64) (*[]CollectionAsset, error) {
	const op errors.Op = "GetCollectionAssetByAsaID"
	query := fmt.Sprintf("select %s from %s.collection_asset where asa_id = ?", strings.Join(CollectionAssetTableKeys(), ","), NFDDatabase())

	var assets []CollectionAsset
	err := h.Select(&assets, query, asaID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Collection Assets Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &assets, nil
}

// DeleteCollectionAssetsNotIn deletes collection assets that are not included in a list for a given collection
func DeleteCollectionAssetsNotIn[H Handle](h H, id string, asaIDs ...interface{}) error {
	const op errors.Op = "DeleteCollectionAssetsNotIn"
	data := append([]interface{}{id}, asaIDs...)
	qMarks := []rune(strings.Repeat("?, ", (len(asaIDs) - 1)))
	query := fmt.Sprintf("delete from %s.collection_asset where id = ? and asa_id not in (%s)", NFDDatabase(), string(qMarks[0:len(qMarks)-2]))
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
