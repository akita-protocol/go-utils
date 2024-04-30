package db

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
)

type StakingPoolStatusType string

const (
	StakingPoolStatusUpcoming StakingPoolStatusType = "upcoming"
	StakingPoolStatusOpen     StakingPoolStatusType = "open"
	StakingPoolStatusClosed   StakingPoolStatusType = "closed"
	StakingPoolStatusEnded    StakingPoolStatusType = "ended"
	StakingPoolStatusUnkown   StakingPoolStatusType = "unknown"
)

type PoolScheduleType string

const (
	Ongoing PoolScheduleType = "ongoing"
	Normal  PoolScheduleType = "normal"
)

type StakingPool struct {
	ID           string                 `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Status       *StakingPoolStatusType `structs:"status,omitempty" db:"status" json:"status,omitempty"`
	Title        string                 `structs:"title,omitempty" db:"title" json:"title,omitempty"`
	About        string                 `structs:"about,omitempty" db:"about" json:"about,omitempty"`
	Avatar       *string                `structs:"avatar,omitempty" db:"avatar" json:"avatar,omitempty"`
	AvatarID     *uint64                `structs:"avatar_id,omitempty" db:"avatar_id" json:"avatar_id,omitempty"`
	Banner       *string                `structs:"banner,omitempty" db:"banner" json:"banner,omitempty"`
	BannerID     *uint64                `structs:"banner_id,omitempty" db:"banner_id" json:"banner_id,omitempty"`
	Slots        *uint64                `structs:"slots,omitempty" db:"slots" json:"slots,omitempty"`
	ScheduleType PoolScheduleType       `structs:"schedule_type,omitempty" db:"schedule_type" json:"schedule_type,omitempty"`
	// Strict entries for staking pools mean that users must sign up during the signup period & if they dont they cant participate
	StrictEntry       *bool      `structs:"strict_entry,omitempty" db:"strict_entry" json:"strict_entry,omitempty"`
	AnnouncementDate  *time.Time `structs:"announcement_date,omitempty" db:"announcement_date" json:"announcement_date,omitempty"`
	SignUpsOpenDate   *time.Time `structs:"signups_open_date,omitempty" db:"signups_open_date" json:"signups_open_date,omitempty"`
	StakingBeginsDate time.Time  `structs:"staking_begins_date,omitempty" db:"staking_begins_date" json:"staking_begins_date,omitempty"`
	StakingEndsDate   *time.Time `structs:"staking_ends_date,omitempty" db:"staking_ends_date" json:"staking_ends_date,omitempty"`
	Creator           string     `structs:"creator,omitempty" db:"creator" json:"creator,omitempty"`
	CreatedDate       time.Time  `structs:"created_date,omitempty" db:"created_date" json:"created_date,omitempty"`
	Address           string     `structs:"address,omitempty" db:"address" json:"address,omitempty"`
}

func (p *StakingPool) GetStakingPoolStatus() string {
	n := time.Now()
	theSignUpWindowHasntOpened := p.SignUpsOpenDate != nil && p.SignUpsOpenDate.After(n)
	theSignUpWindowIsOpen := p.SignUpsOpenDate != nil && n.After(*p.SignUpsOpenDate)
	stakingHasntStarted := n.Before(p.StakingBeginsDate)
	stakingIsEndless := p.StakingEndsDate == nil
	stakingHasStarted := n.After(p.StakingBeginsDate)
	stakingHasntEnded := p.StakingEndsDate == nil || n.Before(*p.StakingEndsDate)
	stakingHasEnded := p.StakingEndsDate != nil && n.After(*p.StakingEndsDate)
	allowLateJoining := p.ScheduleType == "ongoing"

	if theSignUpWindowHasntOpened {
		return "upcoming"
	} else if theSignUpWindowIsOpen && stakingHasntStarted || (allowLateJoining && (stakingIsEndless || stakingHasntEnded)) {
		return "open"
	} else if stakingHasStarted && !allowLateJoining && stakingHasntEnded {
		return "closed"
	} else if stakingHasEnded {
		return "ended"
	}

	return "unknown"
}

var stakingPoolTableKeySlices = map[TableKeySlice][]string{
	TKSTiny:  {"id", "status", "title", "avatar"},
	TKSSmall: {"id", "status", "title", "about", "avatar", "avatar_id", "banner", "banner_id", "slots", "schedule_type", "strict_entry", "announcement_date", "signups_open_date", "staking_begins_date", "staking_ends_date"},
	TKSFull:  {"id", "status", "title", "about", "avatar", "avatar_id", "banner", "banner_id", "slots", "schedule_type", "strict_entry", "announcement_date", "signups_open_date", "staking_begins_date", "staking_ends_date", "creator", "created_date", "address"},
}

func StakingPoolTableKeys(st TableKeySlice) []string {
	return stakingPoolTableKeySlices[st]
}

func GetStakingPool[H Handle](h H, id string, st TableKeySlice) (*StakingPool, error) {
	const op errors.Op = "GetStakingPool"
	query := fmt.Sprintf("select %s from %s.pool where id = ?", strings.Join(StakingPoolTableKeys(st), ","), StakingDatabase())

	var p StakingPool
	err := h.Get(&p, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Pool Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &p, nil
}

func GetStakingPoolByCreator[H Handle](h H, creator string, st TableKeySlice) (*StakingPool, error) {
	const op errors.Op = "GetStakingPoolByCreator"
	query := fmt.Sprintf("select %s from %s.pool where creator = ?", strings.Join(StakingPoolTableKeys(st), ","), StakingDatabase())

	var p StakingPool
	err := h.Get(&p, query, creator)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Pool Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &p, nil
}

func GetStakingPools[H Handle](h H, st TableKeySlice, start uint64, limit uint64) (*[]StakingPool, error) {
	const op errors.Op = "GetStakingPools"
	query := fmt.Sprintf("select %s from %s.pool order by status, title asc limit ?, ?", strings.Join(StakingPoolTableKeys(st), ","), StakingDatabase())

	var pools []StakingPool
	err := h.Select(&pools, query, start, limit)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Pool Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &pools, nil
}

func GetStakingPoolsIn[H Handle](h H, ids []string, st TableKeySlice) (*[]StakingPool, error) {
	const op errors.Op = "GetStakingPoolsIn"
	query := fmt.Sprintf("select %s from %s.pool where id in(%s)", strings.Join(StakingPoolTableKeys(st), ","), StakingDatabase(), strings.Repeat("?, ", len(ids))[0:(len(ids)*3)-2])

	var pools []StakingPool
	err := h.Select(&pools, query, misc.ToInterfaceSlice(ids)...)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}
	return &pools, nil
}

func GetStakingPoolsByCreator[H Handle](h H, creator string, st TableKeySlice, start uint64, limit uint64) (*[]StakingPool, error) {
	const op errors.Op = "GetStakingPoolsByCreator"
	query := fmt.Sprintf("select %s from %s.pool where creator = ? limit ?, ?", strings.Join(StakingPoolTableKeys(st), ","), StakingDatabase())

	var pools []StakingPool
	err := h.Select(&pools, query, creator, start, limit)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}
	return &pools, nil
}

func GetStakingPoolsIDs[H Handle](h H) (*[]string, error) {
	const op errors.Op = "GetStakingPoolsIDs"
	query := fmt.Sprintf("select id from %s.pool order by id asc", StakingDatabase())

	var ids []string
	err := h.Select(&ids, query)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}
	return &ids, nil
}

func GetStakingPoolAddress[H Handle](h H, id string) (string, error) {
	const op errors.Op = "GetStakingPoolAddress"
	query := fmt.Sprintf("select address from %s.pool where id = ?", StakingDatabase())
	var address string
	err := h.Get(&address, query, id)
	if err != nil {
		return "", errors.E(pkg, op, err)
	}
	return address, nil
}
