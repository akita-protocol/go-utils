package db

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
)

type StakingRewardType string

const (
	StakingRewardTypeAsset StakingRewardType = "asset"
	StakingRewardTypeCode  StakingRewardType = "code"
)

type StakingRewardDistributionType string

const (
	StakingRewardDistributionTypeDaily StakingRewardDistributionType = "daily"
	StakingRewardDistributionTypeEnd   StakingRewardDistributionType = "end"
)

type StakingRewardRateType string

const (
	StakingRewardRateTypePercentile StakingRewardRateType = "percentile"
	StakingRewardRateTypeFlat       StakingRewardRateType = "flat"
	StakingRewardRateTypeEven       StakingRewardRateType = "even"
	StakingRewardRateTypeShuffle    StakingRewardRateType = "shuffle"
)

type StakingReward struct {
	ID               string                        `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	PoolID           string                        `structs:"pool_id,omitempty" db:"pool_id" json:"pool_id,omitempty"`
	Type             StakingRewardType             `structs:"type,omitempty" db:"type" json:"type,omitempty"`
	DistributionType StakingRewardDistributionType `structs:"distribution_type,omitempty" db:"distribution_type" json:"distribution_type,omitempty"`
	RateType         StakingRewardRateType         `structs:"rate_type,omitempty" db:"rate_type" json:"rate_type,omitempty"`
	Rate             uint                          `structs:"rate,omitempty" db:"rate" json:"rate,omitempty"`
	Title            string                        `structs:"title,omitempty" db:"title" json:"title,omitempty"`
	About            *string                       `structs:"about,omitempty" db:"about" json:"about,omitempty"`
	StakingRewardURL string                        `structs:"reward_url,omitempty" db:"reward_url" json:"reward_url,omitempty"`

	// reward type asset fields
	AssetID  *uint64 `structs:"asset_id,omitempty" db:"asset_id" json:"asset_id,omitempty"`
	Quantity *uint64 `structs:"quantity,omitempty" db:"quantity" json:"quantity,omitempty"`

	// reward type code fields
	MerchantID         *string        `structs:"merchant_id,omitempty" db:"merchant_id" json:"merchant_id,omitempty"`
	ExpirationDuration *time.Duration `structs:"expiration_duration,omitempty" db:"expiration_duration" json:"expiration_duration,omitempty"`
}

func StakingRewardTableKeys() []string {
	return []string{
		"id",
		"pool_id",
		"type",
		"distribution_type",
		"rate_type",
		"rate",
		"title",
		"about",
		"reward_url",
		"asset_id",
		"quantity",
		"merchant_id",
		"expiration_duration",
	}
}

func GetStakingRewardByID[H Handle](h H, id string) (*StakingReward, error) {
	const op errors.Op = "GetStakingReward"
	query := fmt.Sprintf("select %s from %s.reward where id = ?", strings.Join(StakingRewardTableKeys(), ","), StakingDatabase())

	var r StakingReward
	err := h.Get(&r, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &r, nil
}

func GetStakingRewardsByID[H Handle](h H, ids []string) (*[]StakingReward, error) {
	const op errors.Op = "GetStakingReward"
	query := fmt.Sprintf("select %s from %s.reward where id in(%s)", strings.Join(StakingRewardTableKeys(), ","), StakingDatabase(), strings.Repeat("?, ", len(ids))[0:(len(ids)*3)-2])

	var r []StakingReward
	err := h.Select(&r, query, misc.ToInterfaceSlice(ids)...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &r, nil
}

func GetStakingRewardsByPoolID[H Handle](h H, poolID string) (*[]StakingReward, error) {
	const op errors.Op = "GetStakingRewardsByPoolID"
	query := fmt.Sprintf("select %s from %s.reward where pool_id = ?", strings.Join(StakingRewardTableKeys(), ","), StakingDatabase())

	var list []StakingReward
	err := h.Select(&list, query, poolID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetStakingRewardsNeedingSnapshots[H Handle](h H) (*[]StakingReward, error) {
	const op errors.Op = "GetStakingRewardsNeedingSnapshots"
	query := fmt.Sprintf("select %s from %s.reward where distribution_type = 'daily' and exists(select id from pool where staking_begins_date < ? and id = pool_id)", strings.Join(StakingRewardTableKeys(), ","), StakingDatabase())
	t := time.Now()

	var list []StakingReward
	err := h.Select(&list, query, t.Format(TimeFormat))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetStakingRewardsQualifiedForDisbursement[H Handle](h H) (*[]StakingReward, error) {
	const op errors.Op = "GetStakingRewardsQualifiedForDisbursement"
	query := fmt.Sprintf("select %s from %s.reward where (distribution_type = 'daily' and exists(select id from pool where id = pool_id and (status != 'ended' or status is null) and staking_begins_date < ?)) or (distribution_type = 'end' and exists(select id from pool where id = pool_id and (status != 'ended' or status is null) and schedule_type = 'normal' and staking_ends_date < ? and staking_ends_date > ?))", strings.Join(StakingRewardTableKeys(), ","), StakingDatabase())

	loc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}

	t := time.Now().In(loc)
	hours := 24
	n := t.Add(time.Duration(-hours) * time.Hour)
	// change to within the last day

	var list []StakingReward
	err = h.Select(&list, query, t.Format(TimeFormat), t.Format(TimeFormat), n.Format(TimeFormat))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

// func GetPoolsRequiringDisbursements[H Handle](h H) {
// 	const op errors.Op = "GetPoolsRequiringDisbursements"
// 	query := "select id from pool where staking_ends_date > ? and staking_ends_date < ?"
// 	t := time.Now()
// 	n := t.Add(time.Hour)

// 	var ids []string
// 	err := h.Select(&ids, query, t, n)
// 	if err != nil {
// 		return nil, errors.E(pkg, op, err)
// 	}
// 	return &ids, nil
// }
