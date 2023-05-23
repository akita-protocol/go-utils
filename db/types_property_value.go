package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
)

type PropertyValue struct {
	// ID is the id of the property
	ID                    string  `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Name                  string  `structs:"name,omitempty" db:"name" json:"name,omitempty"`
	Image                 *string `structs:"image,omitempty" db:"image" json:"image,omitempty"`
	ImageIntegrity        *string `structs:"image_integrity,omitempty" db:"image_integrity" json:"image_integrity,omitempty"`
	ImageMimeType         *string `structs:"image_mimetype,omitempty" db:"image_mimetype" json:"image_mimetype,omitempty"`
	AnimationURL          *string `structs:"animation_url,omitempty" db:"animation_url" json:"animation_url,omitempty"`
	AnimationURLIntegrity *string `structs:"animation_url_integrity,omitempty" db:"animation_url_integrity" json:"animation_url_integrity,omitempty"`
	AnimationURLMimeType  *string `structs:"animation_url_mimetype,omitempty" db:"animation_url_mimetype" json:"animation_url_mimetype,omitempty"`
}

func PropertyValueTableKeys() []string {
	return []string{"id", "name", "image", "image_integrity", "image_mime_type", "animation_url", "animation_url_integrity", "animation_url_mime_type"}
}

func GetPropertyValues[H Handle](h H, id string) (*[]PropertyValue, error) {
	const op errors.Op = "GetPropertyValues"
	query := fmt.Sprintf("select %s from %s.property_value where id = ?", strings.Join(PropertyValueTableKeys(), ","), NFDDatabase())

	var propertyValues []PropertyValue
	err := h.Select(&propertyValues, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Property Values Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &propertyValues, nil
}

// DeletePropertyValueNotIn deletes properties that are not included in a list for a given NFD
func DeletePropertyValueNotIn[H Handle](h H, id string, names ...string) error {
	const op errors.Op = "DeletePropertyValueNotIn"
	var err error
	data := append([]interface{}{id}, misc.ToInterfaceSlice(names)...)
	query := fmt.Sprintf("delete from %s.property_value where id = ?", NFDDatabase())

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
