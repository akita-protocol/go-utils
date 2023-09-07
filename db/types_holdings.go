package db

import (
	"database/sql"
	"fmt"
	"math/big"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
)

type Holdings struct {
	Address          string  `structs:"address,omitempty" db:"address" json:"address,omitempty"`
	AsaID            *uint64 `structs:"asa_id,omitempty" db:"asa_id" json:"asa_id"`
	Amount           *uint64 `structs:"amount,omitempty" db:"amount" json:"amount"`
	LastChangedBlock *uint64 `structs:"last_changed_block,omitempty" db:"last_changed_block" json:"last_changed_block,omitempty"`
}

func HoldingsTableKeys() []string {
	return []string{
		"address",
		"asa_id",
		"amount",
		"last_changed_block",
	}
}

func GetAllHoldingWallets[H Handle](h H) (*[]string, error) {
	const op errors.Op = "GetAllHoldingWallets"
	query := fmt.Sprintf("select distinct(address) from %s.holdings", HoldingsDatabase())
	var list []string

	err := h.Select(&list, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Holdings Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &list, nil
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

func GetHoldingsByAsaIDIn[H Handle](h H, ids []uint64) (*[]Holdings, error) {
	const op errors.Op = "GetHoldingsByAsaIDIn"
	qMarks := []rune(strings.Repeat("?, ", len(ids)))
	query := fmt.Sprintf("select address, asa_id, amount from %s.holdings where asa_id in (%s)", HoldingsDatabase(), string(qMarks[0:len(qMarks)-2]))
	var list []Holdings

	err := h.Select(&list, query, misc.ToInterfaceSlice(ids)...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Holdings Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &list, nil
}

func GetHoldingsAddressesByAsaIDIn[H Handle](h H, ids []uint64) (*[]string, error) {
	const op errors.Op = "GetHoldingsAddressesByAsaIDIn"
	qMarks := []rune(strings.Repeat("?, ", len(ids)))
	query := fmt.Sprintf("select distinct(address) from %s.holdings where asa_id in (%s)", HoldingsDatabase(), string(qMarks[0:len(qMarks)-2]))
	var list []string

	err := h.Select(&list, query, misc.ToInterfaceSlice(ids)...)
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

func GetHoldingsSumByAddressListAndAsaID[H Handle](h H, asaID uint64, addresses []string) (*big.Int, error) {
	const op errors.Op = "GetHoldingsSumByAddressListAndAsaID"
	qMarks := []rune(strings.Repeat("?, ", len(addresses)))
	query := fmt.Sprintf("select sum(amount) from %s.holdings where asa_id = ? and address in (%s)", HoldingsDatabase(), string(qMarks[0:len(qMarks)-2]))
	var sumString string

	err := h.Get(&sumString, query, append([]interface{}{asaID}, misc.ToInterfaceSlice(addresses)...)...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Holdings Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	var bignum, ok = new(big.Int).SetString(sumString, 0)
	if !ok {
		return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, fmt.Errorf("cant initialize big int"))
	}

	return bignum, nil
}
