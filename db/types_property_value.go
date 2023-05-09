package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type PropertyValue struct {
	ID                    string  `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Name                  string  `structs:"name,omitempty" db:"name" json:"name,omitempty"`
	Image                 *string `structs:"image,omitempty" db:"image" json:"image,omitempty"`
	ImageIntegrity        *string `structs:"image_integrity,omitempty" db:"image_integrity" json:"image_integrity,omitempty"`
	ImageMimeType         *string `structs:"image_mime_type,omitempty" db:"image_mime_type" json:"image_mime_type,omitempty"`
	AnimationURL          *string `structs:"animation_url,omitempty" db:"animation_url" json:"animation_url,omitempty"`
	AnimationURLIntegrity *string `structs:"animation_url_integrity,omitempty" db:"animation_url_integrity" json:"animation_url_integrity,omitempty"`
	AnimationURLMimeType  *string `structs:"animation_url_mime_type,omitempty" db:"animation_url_mime_type" json:"animation_url_mime_type,omitempty"`
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
