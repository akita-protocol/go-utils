package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type CollectionExtras struct {
	ID    string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Key   string `structs:"key,omitempty" db:"key" json:"key,omitempty"`
	Value string `structs:"value,omitempty" db:"value" json:"value,omitempty"`
}

func CollectionExtrasTableKeys() []string {
	return []string{"id", "key", "value"}
}

func GetCollectionExtras[H Handle](h H, id string) (*[]CollectionExtras, error) {
	const op errors.Op = "GetCollectionExtras"
	query := fmt.Sprintf("select %s from %s.collection_extras where id = ?", strings.Join(CollectionExtrasTableKeys(), ","), NFDDatabase())

	var extras []CollectionExtras
	err := h.Select(&extras, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Collection Extras Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &extras, nil
}

func GetCollectionExtrasWithKey[H Handle](h H, id string, key string) (*CollectionExtras, error) {
	const op errors.Op = "GetCollectionExtrasWithKey"
	query := fmt.Sprintf("select %s from %s.collection_extras where id = ? and key = ?", strings.Join(CollectionExtrasTableKeys(), ","), NFDDatabase())

	var extra CollectionExtras
	err := h.Get(&extra, query, id, key)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Collection Extras Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &extra, nil
}
