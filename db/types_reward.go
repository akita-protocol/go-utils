package db

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
)

type RewardType string

const (
	RewardTypeAsset RewardType = "asset"
	RewardTypeCode  RewardType = "code"
)

type RewardDistributionType string

const (
	RewardDistributionTypeDaily RewardDistributionType = "daily"
	RewardDistributionTypeEnd   RewardDistributionType = "end"
)

type RewardRateType string

const (
	RewardRateTypePercentile RewardRateType = "percentile"
	RewardRateTypeFlat       RewardRateType = "flat"
	RewardRateTypeEven       RewardRateType = "even"
	RewardRateTypeShuffle    RewardRateType = "shuffle"
)

type Reward struct {
	ID               string                 `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	PoolID           string                 `structs:"pool_id,omitempty" db:"pool_id" json:"pool_id,omitempty"`
	Type             RewardType             `structs:"type,omitempty" db:"type" json:"type,omitempty"`
	DistributionType RewardDistributionType `structs:"distribution_type,omitempty" db:"distribution_type" json:"distribution_type,omitempty"`
	RateType         RewardRateType         `structs:"rate_type,omitempty" db:"rate_type" json:"rate_type,omitempty"`
	Rate             uint                   `structs:"rate,omitempty" db:"rate" json:"rate,omitempty"`
	Title            string                 `structs:"title,omitempty" db:"title" json:"title,omitempty"`
	About            *string                `structs:"about,omitempty" db:"about" json:"about,omitempty"`
	RewardURL        string                 `structs:"reward_url,omitempty" db:"reward_url" json:"reward_url,omitempty"`

	// reward type asset fields
	AssetID  *uint64 `structs:"asset_id,omitempty" db:"asset_id" json:"asset_id,omitempty"`
	Quantity *uint64 `structs:"quantity,omitempty" db:"quantity" json:"quantity,omitempty"`

	// reward type code fields
	MerchantID         *string        `structs:"merchant_id,omitempty" db:"merchant_id" json:"merchant_id,omitempty"`
	ExpirationDuration *time.Duration `structs:"expiration_duration,omitempty" db:"expiration_duration" json:"expiration_duration,omitempty"`

	Address *string `structs:"address,omitempty" db:"address" json:"address,omitempty"`
}

func RewardTableKeys() []string {
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
		"address",
	}
}

func GetRewardByID[H Handle](h H, id string) (*Reward, error) {
	const op errors.Op = "GetReward"
	query := fmt.Sprintf("select %s from %s.reward where id = ?", strings.Join(RewardTableKeys(), ","), StakingDatabase())

	var r Reward
	err := h.Get(&r, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &r, nil
}

func GetRewardsByID[H Handle](h H, ids []string) (*[]Reward, error) {
	const op errors.Op = "GetReward"
	query := fmt.Sprintf("select %s from %s.reward where id in(%s)", strings.Join(RewardTableKeys(), ","), StakingDatabase(), strings.Repeat("?, ", len(ids))[0:(len(ids)*3)-2])

	var r []Reward
	err := h.Select(&r, query, misc.ToInterfaceSlice(ids)...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &r, nil
}

func GetRewardsByPoolID[H Handle](h H, poolID string) (*[]Reward, error) {
	const op errors.Op = "GetRewardsByPoolID"
	query := fmt.Sprintf("select %s from %s.reward where pool_id = ?", strings.Join(RewardTableKeys(), ","), StakingDatabase())

	var list []Reward
	err := h.Select(&list, query, poolID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetRewardsNeedingSnapshots[H Handle](h H) (*[]Reward, error) {
	const op errors.Op = "GetRewardsNeedingSnapshots"
	query := fmt.Sprintf("select %s from %s.reward where distribution_type = 'daily' and exists(select id from pool where staking_begins_date < ? and id = pool_id)", strings.Join(RewardTableKeys(), ","), StakingDatabase())
	t := time.Now()

	var list []Reward
	err := h.Select(&list, query, t.Format(TimeFormat))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Reward Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetRewardsQualifiedForDisbursement[H Handle](h H) (*[]Reward, error) {
	const op errors.Op = "GetRewardsQualifiedForDisbursement"
	query := fmt.Sprintf("select %s from %s.reward where (distribution_type = 'daily' and exists(select id from pool where id = pool_id and (status != 'ended' or status is null) and staking_begins_date < ?)) or (distribution_type = 'end' and exists(select id from pool where id = pool_id and (status != 'ended' or status is null) and schedule_type = 'normal' and staking_ends_date < ? and staking_ends_date > ?))", strings.Join(RewardTableKeys(), ","), StakingDatabase())

	loc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}

	t := time.Now().In(loc)
	hours := 24
	n := t.Add(time.Duration(-hours) * time.Hour)
	// change to within the last day

	var list []Reward
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
