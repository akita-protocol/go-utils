package compound

import "gitlab.algotoolchain.com/keelakay/utils/db"

type Property struct {
	*db.Property
	Values []db.PropertyValue       `json:"values,omitempty"`
	Extras []db.PropertyValueExtras `json:"extras,omitempty"`
}
