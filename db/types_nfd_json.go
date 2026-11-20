package db

import (
	"database/sql"
	"fmt"

	"github.com/akita-protocol/go-utils/errors"
)

type NFDJson struct {
	// ID is the app id of the NFD
	ID   uint64 `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Data string `structs:"data,omitempty" db:"data" json:"data,omitempty"`
}

func GetNFDJson[H DBStruct](h H, nfdID string) (*NFDJson, error) {
	const op errors.Op = "GetNFDJson"
	query := fmt.Sprintf("select id, data from %s.nfd_json where id = ?", NFDDatabase())
	var list NFDJson

	err := h.Get(&list, query, nfdID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "NFDJson Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &list, nil
}
