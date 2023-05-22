package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
)

type PropertyValueExtras struct {
	ID    string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Name  string `structs:"name,omitempty" db:"name" json:"name,omitempty"`
	Key   string `structs:"key,omitempty" db:"key" json:"key,omitempty"`
	Value string `structs:"value,omitempty" db:"value" json:"value,omitempty"`
}

func PropertyValueExtrasTableKeys() []string {
	return []string{"id", "name", "'key'", "value"}
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

// DeletePropertyValueExtrasNotIn deletes properties that are not included in a list for a given NFD
func DeletePropertyValueExtrasNotIn[H Handle](h H, id string, name string, keys ...string) error {
	const op errors.Op = "DeletePropertyValueExtrasNotIn"
	var err error
	data := append([]interface{}{id, name}, misc.ToInterfaceSlice(keys)...)
	query := fmt.Sprintf("delete from %s.property_value_extras where id = ? and name = ?", NFDDatabase())

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
