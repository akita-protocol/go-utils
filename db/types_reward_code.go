package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/akita-protocol/go-utils/errors"
)

type RewardCode struct {
	RewardID string  `structs:"reward_id,omitempty" db:"reward_id" json:"reward_id,omitempty"`
	EntryID  *string `structs:"entry_id,omitempty" db:"entry_id" json:"entry_id,omitempty"`
	Code     string  `structs:"code,omitempty" db:"code" json:"code,omitempty"`
	Redeemed bool    `structs:"redeemed,omitempty" db:"redeemed" json:"redeemed,omitempty"`
}

func RewardCodeTableKeys() []string {
	return []string{
		"reward_id",
		"entry_id",
		"code",
		"redeemed",
	}
}

func GetRewardCodesByRewardID[H Handle](h H, rewardID string) (*[]RewardCode, error) {
	const op errors.Op = "GetRewardCodeByRewardID"
	query := fmt.Sprintf("select %s from %s.reward_code where reward_id = ?", strings.Join(RewardCodeTableKeys(), ","), StakingDatabase())

	var list []RewardCode
	err := h.Select(&list, query, rewardID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Codes Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetRewardCodesByEntryID[H Handle](h H, entryID string) (*[]RewardCode, error) {
	const op errors.Op = "GetRewardCodeByEntryID"
	query := fmt.Sprintf("select %s from %s.reward_code where entry_id = ?", strings.Join(RewardCodeTableKeys(), ","), StakingDatabase())

	var list []RewardCode
	err := h.Select(&list, query, entryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Codes Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}
