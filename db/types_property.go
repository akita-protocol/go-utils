package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
)

type Property struct {
	ID   string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Name string `structs:"name,omitempty" db:"name" json:"name,omitempty"`
}

func PropertyTableKeys() []string {
	return []string{"id", "name"}
}

func GetProperties[H Handle](h H, id string) (*[]Property, error) {
	const op errors.Op = "GetProperties"
	query := fmt.Sprintf("select %s from %s.property where id = ?", strings.Join(PropertyTableKeys(), ","), NFDDatabase())

	var properties []Property
	err := h.Select(&properties, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Properties Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &properties, nil
}

func GetPropertiesByName[H Handle](h H, id string, name string) (*[]Property, error) {
	const op errors.Op = "GetProperties"
	query := fmt.Sprintf("select %s from %s.property where id = ? and name = ?", strings.Join(PropertyTableKeys(), ","), NFDDatabase())

	var properties []Property
	err := h.Select(&properties, query, id, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Properties Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &properties, nil
}

// DeletePropertyNotIn deletes properties that are not included in a list for a given NFD
func DeletePropertyNotIn[H Handle](h H, id string, names ...string) error {
	const op errors.Op = "DeletePropertyNotIn"
	var err error
	data := append([]interface{}{id}, misc.ToInterfaceSlice(names)...)
	query := fmt.Sprintf("delete from %s.property where id = ?", NFDDatabase())

	if len(names) > 0 {
		qMarks := []rune(strings.Repeat("?, ", len(names)))
		query += fmt.Sprintf(" and name not in (%s)", string(qMarks[0:len(qMarks)-2]))
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
