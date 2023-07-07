package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

// whitelist token for shuffle
type ShuffleWhitelistToken struct {
	ID              string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	ShuffleID       string `structs:"shuffle_id,omitempty" db:"shuffle_id" json:"shuffle_id,omitempty"`
	AssetID         uint64 `structs:"asset_id,omitempty" db:"asset_id" json:"asset_id,omitempty"`
	ReserveQuantity uint64 `structs:"reserve_quantity,omitempty" db:"reserve_quantity" json:"reserve_quantity,omitempty"`
}

func ShuffleWhitelistTokenTableKeys() []string {
	return []string{
		"id",
		"shuffle_id",
		"asset_id",
		"reserve_quantity",
	}
}

func GetShuffleWhitelistTokens[H Handle](h H, shuffleID string) (*[]ShuffleWhitelistToken, error) {
	const op errors.Op = "GetShuffleWhitelistTokens"
	query := fmt.Sprintf("select %v from %v.shuffle_whitelist_token where shuffle_id = ?", strings.Join(ShuffleWhitelistTokenTableKeys(), ","), ShuffleDatabase())

	var tokens []ShuffleWhitelistToken
	err := h.Select(&tokens, query, shuffleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Shuffle Whitelist Token Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &tokens, err
}

func GetShuffleWhitelistToken[H Handle](h H, id string) (*ShuffleWhitelistToken, error) {
	const op errors.Op = "GetShuffleWhitelistToken"
	query := fmt.Sprintf("select %v from %v.shuffle_whitelist_token where id = ?", strings.Join(ShuffleWhitelistTokenTableKeys(), ","), ShuffleDatabase())

	var token ShuffleWhitelistToken
	err := h.Get(&token, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Shuffle Whitelist Token Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &token, err
}
