package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

// payment options either for the shuffle globally or a specific whitelist token
type ShufflePaymentOption struct {
	ID           string  `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	ShuffleID    string  `structs:"shuffle_id,omitempty" db:"shuffle_id" json:"shuffle_id,omitempty"`
	WhitelistID  *string `structs:"whitelist_id,omitempty" db:"whitelist_id" json:"whitelist_id,omitempty"`
	AssetID      uint64  `structs:"asset_id,omitempty" db:"asset_id" json:"asset_id,omitempty"`
	Amount       *uint64 `structs:"amount,omitempty" db:"amount" json:"amount,omitempty"`
	AmountInAlgo *uint64 `structs:"amount_in_algo,omitempty" db:"amount_in_algo" json:"amount_in_algo,omitempty"`
}

func ShufflePaymentOptionTableKeys() []string {
	return []string{
		"id",
		"shuffle_id",
		"whitelist_id",
		"asset_id",
		"amount",
		"amount_in_algo",
	}
}

func GetShufflePaymentOptions[H Handle](h H, shuffleID string) (*[]ShufflePaymentOption, error) {
	const op errors.Op = "GetShufflePaymentOptions"
	query := fmt.Sprintf("select %v from %v.shuffle_payment_option where shuffle_id = ?", strings.Join(ShufflePaymentOptionTableKeys(), ","), ShuffleDatabase())

	var options []ShufflePaymentOption
	err := h.Select(&options, query, shuffleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Shuffle Payment Option Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &options, err
}
