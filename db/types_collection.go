package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type Collection struct {
	ID          string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	NFDID       string `structs:"nfd_id,omitempty" db:"nfd_id" json:"nfd_id,omitempty"`
	Name        string `structs:"name,omitempty" db:"name" json:"name,omitempty"`
	Description string `structs:"description,omitempty" db:"description" json:"description,omitempty"`
	Banner      uint64 `structs:"banner,omitempty" db:"banner" json:"banner,omitempty"`
	Avatar      uint64 `structs:"avatar,omitempty" db:"avatar" json:"avatar,omitempty"`
	Explicit    bool   `structs:"explicit,omitempty" db:"explicit" json:"explicit,omitempty"`
}

func CollectionTableKeys() []string {
	return []string{"id", "nfd_id", "name", "description", "banner", "avatar", "explicit"}
}

func GetCollection[H Handle](h H, id string) (*Collection, error) {
	const op errors.Op = "GetCollection"
	query := fmt.Sprintf("select %s from %s.collection where id = ?", strings.Join(CollectionTableKeys(), ","), NFDDatabase())

	var c Collection
	err := h.Get(&c, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Collection Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &c, nil
}

func GetCollectionsByNFDID[H Handle](h H, nfdID string) (*[]Collection, error) {
	const op errors.Op = "GetCollectionsPaginated"
	query := fmt.Sprintf("select %s from %s.collection where nfd_id = ?", strings.Join(CollectionTableKeys(), ","), NFDDatabase())

	var collections []Collection
	err := h.Select(&collections, query, nfdID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Collections Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &collections, nil
}

func GetCollectionsPaginated[H Handle](h H, start, limit uint64) (*[]Collection, error) {
	const op errors.Op = "GetCollectionsPaginated"
	query := fmt.Sprintf("select %s from %s.collection limit ?, ?", strings.Join(CollectionTableKeys(), ","), NFDDatabase())

	var collections []Collection
	err := h.Select(&collections, query, start, limit)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Collections Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &collections, nil
}

// DeletCollectionNotIn deletes collections that are not included in a list for a given NFD
func DeletCollectionNotIn[H Handle](h H, nfdID string, ids ...interface{}) error {
	const op errors.Op = "DeletCollectionNotIn"
	data := append([]interface{}{nfdID}, ids...)
	qMarks := []rune(strings.Repeat("?, ", (len(ids) - 1)))
	query := "delete from %s.collection where nfd_id = ? and id not in (" + string(qMarks[0:len(qMarks)-2]) + ")"
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
