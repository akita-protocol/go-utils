package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type CollectionArtist struct {
	ID      string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Address string `structs:"address,omitempty" db:"address" json:"address,omitempty"`
}

func CollectionArtistTableKeys() []string {
	return []string{"id", "address"}
}

func GetCollectionArtist[H Handle](h H) ([]*CollectionArtist, error) {
	const op errors.Op = "GetCollectionArtist"
	query := fmt.Sprintf("select %s from %s.collection_creator", strings.Join(CollectionArtistTableKeys(), ","), NFDDatabase())

	var ccs []*CollectionArtist
	err := h.Select(&ccs, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "CollectionArtist Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return ccs, nil
}

func GetCollectionArtistByCollection[H Handle](h H, collectionID string) ([]*CollectionArtist, error) {
	const op errors.Op = "GetCollectionArtistByCollection"
	query := fmt.Sprintf("select %s from %s.collection_creator where collection_id = ?", strings.Join(CollectionArtistTableKeys(), ","), NFDDatabase())

	var ccs []*CollectionArtist
	err := h.Select(&ccs, query, collectionID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "CollectionArtist Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return ccs, nil
}
