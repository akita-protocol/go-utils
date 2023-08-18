package db

import (
	"database/sql"
	"fmt"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type Holdings struct {
	Address string  `structs:"address,omitempty" db:"address" json:"address,omitempty"`
	AsaID   uint64  `structs:"asa_id,omitempty" db:"asa_id" json:"asa_id"`
	Amount  *uint64 `structs:"amount,omitempty" db:"amount" json:"amount"`
}

func HoldingsTableKeys() []string {
	return []string{
		"address",
		"asa_id",
		"amount",
	}
}

func GetHoldingsByAddress[H Handle](h H, address string) (*[]Holdings, error) {
	const op errors.Op = "GetHoldingsByAddress"
	query := fmt.Sprintf("select address, asa_id, amount from %s.holdings where address = ?", HoldingsDatabase())
	var list []Holdings

	err := h.Select(&list, query, address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Holdings Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &list, nil
}

func GetHoldingsByAsaID[H Handle](h H, asaID uint64) (*[]Holdings, error) {
	const op errors.Op = "GetHoldingsByAsaID"
	query := fmt.Sprintf("select address, asa_id, amount from %s.holdings where asa_id = ?", HoldingsDatabase())
	var list []Holdings

	err := h.Select(&list, query, asaID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Holdings Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &list, nil
}

func GetHoldingsByAddressAndAsaID[H Handle](h H, address string, asaID uint64) (*Holdings, error) {
	const op errors.Op = "GetHoldingsByAddressAndAsaID"
	query := fmt.Sprintf("select address, asa_id, amount from %s.holdings where address = ? and asa_id = ?", HoldingsDatabase())
	var list Holdings

	err := h.Get(&list, query, address, asaID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Holdings Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &list, nil
}
