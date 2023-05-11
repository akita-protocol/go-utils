package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
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

// DeleteCollectionPrefixesNotIn deletes collection prefixes that are not included in a list for a given collection
func DeleteCollectionPrefixesNotIn[H Handle](h H, id string, prefixes ...interface{}) error {
	const op errors.Op = "DeleteCollectionPrefixesNotIn"
	var err error
	data := append([]interface{}{id}, prefixes...)
	query := fmt.Sprintf("delete from %s.collection_prefix where id = ?", NFDDatabase())

	if len(prefixes) > 0 {
		qMarks := []rune(strings.Repeat("?, ", len(prefixes)))
		query += fmt.Sprintf(" and prefix not in (%s)", string(qMarks[0:len(qMarks)-2]))
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
