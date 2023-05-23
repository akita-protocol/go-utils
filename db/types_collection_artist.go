package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
)

type CollectionArtist struct {
	ID      string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Address string `structs:"address,omitempty" db:"address" json:"address,omitempty"`
}

func CollectionArtistTableKeys() []string {
	return []string{"id", "address"}
}

func GetCollectionArtist[H Handle](h H) (*[]CollectionArtist, error) {
	const op errors.Op = "GetCollectionArtist"
	query := fmt.Sprintf("select %s from %s.collection_artist", strings.Join(CollectionArtistTableKeys(), ","), NFDDatabase())

	var ccs []CollectionArtist
	err := h.Select(&ccs, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "CollectionArtist Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &ccs, nil
}

func GetCollectionArtistByCollection[H Handle](h H, id string) (*[]CollectionArtist, error) {
	const op errors.Op = "GetCollectionArtistByCollection"
	query := fmt.Sprintf("select %s from %s.collection_artist where id = ?", strings.Join(CollectionArtistTableKeys(), ","), NFDDatabase())

	var ccs []CollectionArtist
	err := h.Select(&ccs, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "CollectionArtist Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &ccs, nil
}

// DeleteCollectionArtistsNotIn deletes collection artists that are not included in a list for a given collection
func DeleteCollectionArtistsNotIn[H Handle](h H, id string, addresses ...string) error {
	const op errors.Op = "DeleteCollectionArtistsNotIn"
	var err error
	data := append([]interface{}{id}, misc.ToInterfaceSlice(addresses)...)
	query := fmt.Sprintf("delete from %s.collection_artist where id = ?", NFDDatabase())

	if len(addresses) > 0 {
		qMarks := []rune(strings.Repeat("?, ", len(addresses)))
		query += fmt.Sprintf(" and address not in (%s)", string(qMarks[0:len(qMarks)-2]))
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
