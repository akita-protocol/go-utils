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

// DeletCollectionPrefixesNotIn deletes collection prefixes that are not included in a list for a given collection
func DeletCollectionPrefixesNotIn[H Handle](h H, id string, prefixes ...interface{}) error {
	const op errors.Op = "DeletCollectionPrefixesNotIn"
	data := append([]interface{}{id}, prefixes...)
	qMarks := []rune(strings.Repeat("?, ", (len(prefixes) - 1)))
	query := "delete from %s.collection_prefix where id = ? and prefix not in (" + string(qMarks[0:len(qMarks)-2]) + ")"
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
