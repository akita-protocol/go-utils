package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
)

type CollectionExtras struct {
	ID    string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Key   string `structs:"mkey,omitempty" db:"mkey" json:"key,omitempty"`
	Value string `structs:"mvalue,omitempty" db:"mvalue" json:"value,omitempty"`
}

func CollectionExtrasTableKeys() []string {
	return []string{"id", "mkey", "mvalue"}
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

// DeleteCollectionExtrasNotIn deletes collection extras that are not included in a list for a given collection
func DeleteCollectionExtrasNotIn[H Handle](h H, id string, keys ...string) error {
	const op errors.Op = "DeleteCollectionExtrasNotIn"
	var err error
	data := append([]interface{}{id}, misc.ToInterfaceSlice(keys)...)
	query := fmt.Sprintf("delete from %s.collection_extras where id = ?", NFDDatabase())

	if len(keys) > 0 {
		qMarks := []rune(strings.Repeat("?, ", len(keys)))
		query += fmt.Sprintf(" and key not in (%s)", string(qMarks[0:len(qMarks)-2]))
	}

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
