package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type Property struct {
	ID           string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	CollectionID string `structs:"collection_id,omitempty" db:"collection_id" json:"collection_id,omitempty"`
	Name         string `structs:"name,omitempty" db:"name" json:"name,omitempty"`
}

func PropertyTableKeys() []string {
	return []string{"id", "collection_id", "name"}
}

func GetProperties[H Handle](h H, collectionID string) (*[]Property, error) {
	const op errors.Op = "GetProperties"
	query := fmt.Sprintf("select %s from %s.property where collection_id = ?", strings.Join(PropertyTableKeys(), ","), NFDDatabase())

	var properties []Property
	err := h.Select(&properties, query, collectionID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Properties Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &properties, nil
}

func GetPropertiesByName[H Handle](h H, collectionID string, name string) (*[]Property, error) {
	const op errors.Op = "GetProperties"
	query := fmt.Sprintf("select %s from %s.property where collection_id = ? and name = ?", strings.Join(PropertyTableKeys(), ","), NFDDatabase())

	var properties []Property
	err := h.Select(&properties, query, collectionID, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Properties Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &properties, nil
}

// DeletePropertyNotIn deletes properties that are not included in a list for a given NFD
func DeletePropertyNotIn[H Handle](h H, collectionID string, ids ...interface{}) error {
	const op errors.Op = "DeletePropertyNotIn"
	data := append([]interface{}{collectionID}, ids...)
	qMarks := []rune(strings.Repeat("?, ", (len(ids) - 1)))
	query := fmt.Sprintf("delete from %s.property where collection_id = ? and id not in (%s)", NFDDatabase(), string(qMarks[0:len(qMarks)-2]))
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
