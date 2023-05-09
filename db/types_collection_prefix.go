package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type CollectionPrefix struct {
	ID     string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Prefix string `structs:"prefix,omitempty" db:"prefix" json:"prefix,omitempty"`
}

func CollectionPrefixTableKeys() []string {
	return []string{"id", "prefix"}
}

func GetCollectionPrefixes[H Handle](h H, id string) (*[]CollectionPrefix, error) {
	const op errors.Op = "GetCollectionPrefixes"
	query := fmt.Sprintf("select %s from %s.collection_prefix where id = ?", strings.Join(CollectionPrefixTableKeys(), ","), NFDDatabase())

	var prefixes []CollectionPrefix
	err := h.Select(&prefixes, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Collection Prefixes Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &prefixes, nil
}
