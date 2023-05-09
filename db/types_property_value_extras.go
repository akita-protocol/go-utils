package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type PropertyValueExtras struct {
	ID    string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Name  string `structs:"name,omitempty" db:"name" json:"name,omitempty"`
	Key   string `structs:"key,omitempty" db:"key" json:"key,omitempty"`
	Value string `structs:"value,omitempty" db:"value" json:"value,omitempty"`
}

func PropertyValueExtrasTableKeys() []string {
	return []string{"id", "name", "key", "value"}
}

func GetPropertyValueExtras[H Handle](h H, id string) (*[]PropertyValueExtras, error) {
	const op errors.Op = "GetPropertyValueExtras"
	query := fmt.Sprintf("select %s from %s.property_value_extras where id = ?", strings.Join(PropertyValueExtrasTableKeys(), ","), NFDDatabase())

	var extras []PropertyValueExtras
	err := h.Select(&extras, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Property Value Extras Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &extras, nil
}

// DeletPropertyValueExtrasNotIn deletes properties that are not included in a list for a given NFD
func DeletPropertyValueExtrasNotIn[H Handle](h H, id string, name string, keys ...interface{}) error {
	const op errors.Op = "DeletPropertyValueExtrasNotIn"
	data := append([]interface{}{id, name}, keys...)
	qMarks := []rune(strings.Repeat("?, ", (len(keys) - 1)))
	query := "delete from %s.property_value where id = ? and name = ? and key not in (" + string(qMarks[0:len(qMarks)-2]) + ")"
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
