package db

import (
	"database/sql"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type SubscriptionPass struct {
	ID      uint64 `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Address string `structs:"address,omitempty" db:"address" json:"address,omitempty"`
}

func GetSubscriptionPassesBySubscriptionID[H Handle](h H, id uint64, st TableKeySlice) ([]*SubscriptionPass, error) {
	const op errors.Op = "GetSubscriptionPassesBySubscriptionID"
	const query = "select id, address from subscription_passes where subscription_id = ?"

	var passes []*SubscriptionPass
	err := h.Select(&passes, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Subscription Pass Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return passes, nil
}
