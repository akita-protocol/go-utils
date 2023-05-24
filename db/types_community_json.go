package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type CommunityJson struct {
	ID        uint64 `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Data      string `structs:"data,omitempty" db:"data" json:"data,omitempty"`
	Malformed *bool  `structs:"malformed,omitempty" db:"malformed" json:"malformed,omitempty"`
}

func CommunityJsonTableKeys() []string {
	return []string{"id", "data", "malformed"}
}

func GetCommunityJson[H Handle](h H, id uint64) (*CommunityJson, error) {
	const op errors.Op = "GetCollectionJson"
	query := fmt.Sprintf("select %s from %s.community_json where id = ?", strings.Join(CommunityJsonTableKeys(), ","), NFDDatabase())

	var json CommunityJson
	err := h.Get(&json, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Collection Json Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &json, nil
}
