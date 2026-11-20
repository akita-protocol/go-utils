package db

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/akita-protocol/go-utils/errors"
)

type Disbursement struct {
	ID        string    `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	GroupID   string    `structs:"group_id,omitempty" db:"group_id" json:"group_id,omitempty"`
	RewardID  string    `structs:"reward_id,omitempty" db:"reward_id" json:"reward_id,omitempty"`
	UserID    string    `structs:"user_id,omitempty" db:"user_id" json:"user_id,omitempty"`
	EntryID   *string   `structs:"entry_id,omitempty" db:"entry_id" json:"entry_id,omitempty"`
	Quantity  *uint64   `structs:"quantity,omitempty" db:"quantity" json:"quantity,omitempty"`
	Redeemed  bool      `structs:"redeemed,omitempty" db:"redeemed" json:"redeemed,omitempty"`
	Txn       *string   `structs:"txn,omitempty" db:"txn" json:"txn,omitempty"`
	Timestamp time.Time `structs:"ts,omitempty" db:"ts" json:"timestamp,omitempty"`
}

func DisbursementTableKeys() []string {
	return []string{
		"id",
		"group_id",
		"reward_id",
		"user_id",
		"entry_id",
		"quantity",
		"redeemed",
		"txn",
		"ts",
	}
}

func GetDisbursementsByRewardID[H Handle](h H, rewardID string) (*[]Disbursement, error) {
	const op errors.Op = "GetDisbursementsByRewardID"
	query := fmt.Sprintf("select %s from disbursement where reward_id = ?", strings.Join(DisbursementTableKeys(), ","))

	var list []Disbursement
	err := h.Select(&list, query, rewardID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Disbursements Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetUnclaimedDisbursmentsQuantityAllocatedByRewardID[H Handle](h H, rewardID string) (uint64, error) {
	const op errors.Op = "GetUnclaimedDisbursmentsQuantityAllocatedByRewardID"
	const query = "select coalesce(sum(quantity), 0) from disbursement where reward_id = ? and txn is null"
	var quantity uint64
	err := h.Get(&quantity, query, rewardID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, errors.E(pkg, op, err)
	}
	return quantity, nil
}

func GetDisbursementsByEntryID[H Handle](h H, entryID string) (*[]Disbursement, error) {
	const op errors.Op = "GetDisbursementsByEntryID"
	query := fmt.Sprintf("select %s from disbursement where entry_id = ?", strings.Join(DisbursementTableKeys(), ","))

	var list []Disbursement
	err := h.Select(&list, query, entryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Disbursements Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetUnclaimedDisbursementsByEntryID[H Handle](h H, entryID string) (*[]Disbursement, error) {
	const op errors.Op = "GetUnclaimedDisbursementsByEntryID"
	query := fmt.Sprintf("select %s from disbursement where entry_id = ? and redeemed = 0", strings.Join(DisbursementTableKeys(), ","))

	var list []Disbursement
	err := h.Select(&list, query, entryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Disbursements Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetDisbursementsByUserID[H Handle](h H, userID string) (*[]Disbursement, error) {
	const op errors.Op = "GetDisbursementsByUserID"
	query := fmt.Sprintf("select %s from disbursement where user_id = ?", strings.Join(DisbursementTableKeys(), ","))

	var list []Disbursement
	err := h.Select(&list, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Disbursements Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetUnclaimedDisbursementsByUserID[H Handle](h H, userID string) (*[]Disbursement, error) {
	const op errors.Op = "GetDisbursementsByUserID"
	query := fmt.Sprintf("select %s from disbursement where redeemed = 0 and user_id = ?", strings.Join(DisbursementTableKeys(), ","))

	var list []Disbursement
	err := h.Select(&list, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Disbursements Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetUnclaimedDisbursementsByRewardAndUserID[H Handle](h H, rewardID string, userID string) (*[]Disbursement, error) {
	const op errors.Op = "GetDisbursementsByUserID"
	query := fmt.Sprintf("select %s from disbursement where redeemed = 0 and reward_id = ? and user_id = ?", strings.Join(DisbursementTableKeys(), ","))

	var list []Disbursement
	err := h.Select(&list, query, rewardID, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Disbursements Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetDisbursementsByUserAddress[H Handle](h H, address string) (*[]Disbursement, error) {
	const op errors.Op = "GetDisbursementsByUserID"
	query := fmt.Sprintf("select %s from disbursement where user_id = (select id from user where address = ?)", strings.Join(DisbursementTableKeys(), ","))

	var list []Disbursement
	err := h.Select(&list, query, address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Disbursements Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetUnclaimedDisbursementsByUserAddress[H Handle](h H, address string) (*[]Disbursement, error) {
	const op errors.Op = "GetDisbursementsByUserID"
	query := fmt.Sprintf("select %s from disbursement where redeemed = 0 and user_id = (select id from user where address = ?)", strings.Join(DisbursementTableKeys(), ","))

	var list []Disbursement
	err := h.Select(&list, query, address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Disbursements Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}
