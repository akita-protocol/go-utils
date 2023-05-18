package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type StakingRewardCode struct {
	StakingRewardID string  `structs:"reward_id,omitempty" db:"reward_id" json:"reward_id,omitempty"`
	EntryID         *string `structs:"entry_id,omitempty" db:"entry_id" json:"entry_id,omitempty"`
	Code            string  `structs:"code,omitempty" db:"code" json:"code,omitempty"`
	Redeemed        bool    `structs:"redeemed,omitempty" db:"redeemed" json:"redeemed,omitempty"`
}

func StakingRewardCodeTableKeys() []string {
	return []string{
		"reward_id",
		"entry_id",
		"code",
		"redeemed",
	}
}

func GetStakingRewardCodesByStakingRewardID[H Handle](h H, rewardID string) (*[]StakingRewardCode, error) {
	const op errors.Op = "GetStakingRewardCodeByStakingRewardID"
	query := fmt.Sprintf("select %s from %s.reward_code where reward_id = ?", strings.Join(StakingRewardCodeTableKeys(), ","), StakingDatabase())

	var list []StakingRewardCode
	err := h.Select(&list, query, rewardID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Codes Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetStakingRewardCodesByEntryID[H Handle](h H, entryID string) (*[]StakingRewardCode, error) {
	const op errors.Op = "GetStakingRewardCodeByEntryID"
	query := fmt.Sprintf("select %s from %s.reward_code where entry_id = ?", strings.Join(StakingRewardCodeTableKeys(), ","), StakingDatabase())

	var list []StakingRewardCode
	err := h.Select(&list, query, entryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Codes Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}
