package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type SubscriptionOffering struct {
	Recipient string `structs:"recipient,omitempty" db:"recipient" json:"recipient,omitempty"`
	BoxIndex  uint64 `structs:"box_index,omitempty" db:"box_index" json:"box_index,omitempty"`
	Interval  uint64 `structs:"interval,omitempty" db:"interval" json:"interval,omitempty"`
	ASA       uint64 `structs:"asa,omitempty" db:"asa" json:"asa,omitempty"`
	Amount    uint64 `structs:"amount,omitempty" db:"amount" json:"amount,omitempty"`
	PassCount uint64 `structs:"pass_count,omitempty" db:"pass_count" json:"pass_count,omitempty"`
	CID       string `structs:"cid,omitempty" db:"cid" json:"cid,omitempty"`
	Active    bool   `structs:"active,omitempty" db:"active" json:"active,omitempty"`
	Shutdown  bool   `structs:"shutdown,omitempty" db:"shutdown" json:"shutdown,omitempty"`
}

func SubscriptionOfferingTableKeys(st TableKeySlice) []string {
	return []string{"recipient", "box_index", "interval", "asa", "amount", "pass_count", "cid", "active", "shutdown"}
}

func GetSubscriptionOfferingByRecipientAndBoxIndex[H Handle](h H, recipient string, boxIndex uint64, st TableKeySlice) (*SubscriptionOffering, error) {
	const op errors.Op = "GetSubscriptionOfferingByRecipientAndBoxIndex"
	query := fmt.Sprintf("select %s from subscription_offering where recipient = ? and box_index = ?", strings.Join(SubscriptionOfferingTableKeys(st), ","))

	var offering SubscriptionOffering
	err := h.Get(&offering, query, recipient, boxIndex)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Subscription Offering Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &offering, nil
}

func GetSubscriptionOfferingsByRecipient[H Handle](h H, recipient string, st TableKeySlice) ([]*SubscriptionOffering, error) {
	const op errors.Op = "GetSubscriptionOfferingsByRecipient"
	query := fmt.Sprintf("select %s from subscription_offering where recipient = ?", strings.Join(SubscriptionOfferingTableKeys(st), ","))

	var offerings []*SubscriptionOffering
	err := h.Select(&offerings, query, recipient)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Subscription Offerings Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return offerings, nil
}
