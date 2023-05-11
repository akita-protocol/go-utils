package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type CollectionJson struct {
	ID        string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Data      string `structs:"data,omitempty" db:"data" json:"data,omitempty"`
	Malformed *bool  `structs:"malformed,omitempty" db:"malformed" json:"malformed,omitempty"`
}

func CollectionJsonTableKeys() []string {
	return []string{"id", "data", "malformed"}
}

func GetCollectionJson[H Handle](h H, id string) (*CollectionJson, error) {
	const op errors.Op = "GetCollectionJson"
	query := fmt.Sprintf("select %s from %s.collection_json where id = ?", strings.Join(CollectionJsonTableKeys(), ","), NFDDatabase())

	var json CollectionJson
	err := h.Get(&json, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Collection Json Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &json, nil
}
