package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/akita-protocol/go-utils/errors"
)

type SubscriptionPayment struct {
	ID    string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	SubID uint64 `structs:"sub_id,omitempty" db:"sub_id" json:"sub_id,omitempty"`
}

func SubscriptionPaymentTableKeys() []string {
	return []string{"id", "sub_id"}
}

func GetSubscriptionPaymentByTxID[H Handle](h H, txID string) (*SubscriptionPayment, error) {
	const op errors.Op = "GetSubscriptionPaymentByTxID"
	query := fmt.Sprintf("select %s from %s.subscription_payment where id = ?", strings.Join(SubscriptionPaymentTableKeys(), ","), SubscriptionDatabase())

	var sp SubscriptionPayment
	err := h.Get(&sp, query, txID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Subscription Payment Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &sp, nil
}

func GetSubscriptionPaymentsBySubID[H Handle](h H, subID uint64) ([]*SubscriptionPayment, error) {
	const op errors.Op = "GetSubscriptionPaymentsBySubID"
	query := fmt.Sprintf("select %s from %s.subscription_payment where sub_id = ?", strings.Join(SubscriptionPaymentTableKeys(), ","), SubscriptionDatabase())

	var list []*SubscriptionPayment
	err := h.Select(&list, query, subID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Subscription Payment Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return list, nil
}
