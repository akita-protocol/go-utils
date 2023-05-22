package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
)

type CommunityExtras struct {
	ID    uint64 `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Key   string `structs:"key,omitempty" db:"key" json:"key,omitempty"`
	Value string `structs:"value,omitempty" db:"value" json:"value,omitempty"`
}

func CommunityExtrasTableKeys() []string {
	return []string{"id", "'key'", "value"}
}

func GetCommunityExtras[H Handle](h H, id uint64) (*[]CommunityExtras, error) {
	const op errors.Op = "GetCommunityExtras"
	query := fmt.Sprintf("select %s from %s.community_extras where id = ?", strings.Join(CommunityExtrasTableKeys(), ","), NFDDatabase())

	var extras []CommunityExtras
	err := h.Select(&extras, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Collection Extras Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &extras, nil
}

func GetCommunityExtrasWithKey[H Handle](h H, id uint64, key string) (*CommunityExtras, error) {
	const op errors.Op = "GetCommunityExtrasWithKey"
	query := fmt.Sprintf("select %s from %s.community_extras where id = ? and key = ?", strings.Join(CommunityExtrasTableKeys(), ","), NFDDatabase())

	var extra CommunityExtras
	err := h.Get(&extra, query, id, key)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Collection Extras Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &extra, nil
}

// DeleteCommunityExtrasNotIn deletes collection extras that are not included in a list for a given collection
func DeleteCommunityExtrasNotIn[H Handle](h H, id uint64, keys ...string) error {
	const op errors.Op = "DeleteCommunityExtrasNotIn"
	var err error
	data := append([]interface{}{id}, misc.ToInterfaceSlice(keys)...)
	query := fmt.Sprintf("delete from %s.community_extras where id = ?", NFDDatabase())

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
