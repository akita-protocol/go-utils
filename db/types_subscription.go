package db

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type Subscription struct {
	ID               uint64     `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	StartDate        time.Time  `structs:"start_date,omitempty" db:"start_date" json:"start_date,omitempty"`
	BoxIndex         *uint64    `structs:"box_index,omitempty" db:"box_index" json:"box_index,omitempty"`
	Interval         uint64     `structs:"intervl,omitempty" db:"intervl" json:"interval,omitempty"`
	LastPayment      time.Time  `structs:"last_payment,omitempty" db:"last_payment" json:"last_payment,omitempty"`
	Subscriber       string     `structs:"subscriber,omitempty" db:"subscriber" json:"subscriber,omitempty"`
	Recipient        string     `structs:"recipient,omitempty" db:"recipient" json:"recipient,omitempty"`
	ASA              *uint64    `structs:"asa,omitempty" db:"asa" json:"asa,omitempty"`
	Amount           uint64     `structs:"amount,omitempty" db:"amount" json:"amount,omitempty"`
	PassCount        *uint64    `structs:"pass_count,omitempty" db:"pass_count" json:"pass_count,omitempty"`
	LPA              *time.Time `structs:"lpa,omitempty" db:"lpa" json:"lpa,omitempty"`
	CurrentAsOfBlock uint64     `structs:"current_as_of_block,omitempty" db:"current_as_of_block" json:"current_as_of_block,omitempty"`
	Deleted          *bool      `structs:"deleted,omitempty" db:"deleted" json:"deleted,omitempty"`
	DeletedDate      *time.Time `structs:"deleted_date,omitempty" db:"deleted_date" json:"deleted_date,omitempty"`
}

var SubscriptionTableKeySlices = map[TableKeySlice][]string{
	Tiny:  {"id", "box_index", "intervl", "subscriber", "recipient", "asa", "amount"},
	Small: {"id", "start_date", "box_index", "intervl", "last_payment", "subscriber", "recipient", "asa", "amount", "pass_count", "lpa", "current_as_of_block"},
	Full:  {"id", "start_date", "box_index", "intervl", "last_payment", "subscriber", "recipient", "asa", "amount", "pass_count", "lpa", "current_as_of_block", "deleted", "deleted_date"},
}

func SubscriptionTableKeys(st TableKeySlice) []string {
	return SubscriptionTableKeySlices[st]
}

func GetLeastSubscriptionUpdateBlock[H Handle](h H) (uint64, error) {
	const op errors.Op = "GetLatestUpdateBlock"
	const query = "select min(current_as_of_block) from subscription where deleted = 0"

	var block uint64
	err := h.Get(&block, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Subscription Not Found")
		}
		return 0, errors.E(pkg, op, err)
	}
	return block, nil
}

func GetSubscriptionByID[H Handle](h H, st TableKeySlice) (*Subscription, error) {
	const op errors.Op = "GetSubscriptionByID"
	query := fmt.Sprintf("select %s from subscription where id = ?", strings.Join(SubscriptionTableKeys(st), ","))

	var s Subscription
	err := h.Get(&s, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Subscription Pass Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &s, nil
}

func GetSubscriptionsByRecipient[H Handle](h H, recipient string, st TableKeySlice) ([]*Subscription, error) {
	const op errors.Op = "GetSubscriptionsByRecipient"
	query := fmt.Sprintf("select %s from subscription where recipient = ?", strings.Join(SubscriptionTableKeys(st), ","))
	var list []*Subscription

	err := h.Select(&list, query, recipient)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Subscription Pass Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return list, nil
}

func GetSubscriptionsByRecipientOffering[H Handle](h H, recipient string, boxIndex uint64, st TableKeySlice) ([]*Subscription, error) {
	const op errors.Op = "GetSubscriptionsByRecipientOffering"
	query := fmt.Sprintf("select %s from subscription where recipient = ? and box_index = ?", strings.Join(SubscriptionTableKeys(st), ","))
	var list []*Subscription

	err := h.Select(&list, query, recipient, boxIndex)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Subscription Pass Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return list, nil
}

func GetSubscriptionsBySubscriber[H Handle](h H, subscriber string, st TableKeySlice) ([]*Subscription, error) {
	const op errors.Op = "GetSubscriptionsBySubscriber"
	query := fmt.Sprintf("select %s from subscription where subscriber = ?", strings.Join(SubscriptionTableKeys(st), ","))
	var list []*Subscription

	err := h.Select(&list, query, subscriber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Subscription Pass Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return list, nil
}
