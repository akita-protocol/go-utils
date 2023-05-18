package db

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type StakingDisbursement struct {
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

func StakingDisbursementTableKeys() []string {
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

func GetStakingDisbursementsByRewardID[H Handle](h H, rewardID string) (*[]StakingDisbursement, error) {
	const op errors.Op = "GetStakingDisbursementsByRewardID"
	query := fmt.Sprintf("select %s from disbursement where reward_id = ?", strings.Join(StakingDisbursementTableKeys(), ","))

	var list []StakingDisbursement
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

func GetStakingDisbursementsByEntryID[H Handle](h H, entryID string) (*[]StakingDisbursement, error) {
	const op errors.Op = "GetStakingDisbursementsByEntryID"
	query := fmt.Sprintf("select %s from disbursement where entry_id = ?", strings.Join(StakingDisbursementTableKeys(), ","))

	var list []StakingDisbursement
	err := h.Select(&list, query, entryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Disbursements Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetUnclaimedStakingDisbursementsByEntryID[H Handle](h H, entryID string) (*[]StakingDisbursement, error) {
	const op errors.Op = "GetUnclaimedStakingDisbursementsByEntryID"
	query := fmt.Sprintf("select %s from disbursement where entry_id = ? and redeemed = 0", strings.Join(StakingDisbursementTableKeys(), ","))

	var list []StakingDisbursement
	err := h.Select(&list, query, entryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Disbursements Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetStakingDisbursementsByUserID[H Handle](h H, userID string) (*[]StakingDisbursement, error) {
	const op errors.Op = "GetStakingDisbursementsByUserID"
	query := fmt.Sprintf("select %s from disbursement where user_id = ?", strings.Join(StakingDisbursementTableKeys(), ","))

	var list []StakingDisbursement
	err := h.Select(&list, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Disbursements Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetUnclaimedStakingDisbursementsByUserID[H Handle](h H, userID string) (*[]StakingDisbursement, error) {
	const op errors.Op = "GetStakingDisbursementsByUserID"
	query := fmt.Sprintf("select %s from disbursement where redeemed = 0 and user_id = ?", strings.Join(StakingDisbursementTableKeys(), ","))

	var list []StakingDisbursement
	err := h.Select(&list, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Disbursements Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetUnclaimedStakingDisbursementsByRewardAndUserID[H Handle](h H, rewardID string, userID string) (*[]StakingDisbursement, error) {
	const op errors.Op = "GetStakingDisbursementsByUserID"
	query := fmt.Sprintf("select %s from disbursement where redeemed = 0 and reward_id = ? and user_id = ?", strings.Join(StakingDisbursementTableKeys(), ","))

	var list []StakingDisbursement
	err := h.Select(&list, query, rewardID, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Disbursements Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetStakingDisbursementsByUserAddress[H Handle](h H, address string) (*[]StakingDisbursement, error) {
	const op errors.Op = "GetStakingDisbursementsByUserID"
	query := fmt.Sprintf("select %s from disbursement where user_id = (select id from user where address = ?)", strings.Join(StakingDisbursementTableKeys(), ","))

	var list []StakingDisbursement
	err := h.Select(&list, query, address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Disbursements Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetUnclaimedStakingDisbursementsByUserAddress[H Handle](h H, address string) (*[]StakingDisbursement, error) {
	const op errors.Op = "GetStakingDisbursementsByUserID"
	query := fmt.Sprintf("select %s from disbursement where redeemed = 0 and user_id = (select id from user where address = ?)", strings.Join(StakingDisbursementTableKeys(), ","))

	var list []StakingDisbursement
	err := h.Select(&list, query, address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Disbursements Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}
