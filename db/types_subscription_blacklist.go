package db

import (
	"database/sql"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type SubscriptionBlacklist struct {
	Recipient string `structs:"recipient,omitempty" db:"recipient" json:"recipient,omitempty"`
	Address   string `structs:"address,omitempty" db:"address" json:"address,omitempty"`
}

func GetSubscriptionBlacklistByRecipient[H Handle](h H, recipient string, st TableKeySlice) ([]*SubscriptionBlacklist, error) {
	const op errors.Op = "GetSubscriptionBlacklistByRecipient"
	const query = "select recipient, address from subscription_blacklist where recipient = ?"

	var blacklist []*SubscriptionBlacklist
	err := h.Select(&blacklist, query, recipient)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Subscription Blacklist Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return blacklist, nil
}
