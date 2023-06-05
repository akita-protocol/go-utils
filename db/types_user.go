package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
)

type User struct {
	ID        string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	DiscordID string `structs:"discord_id,omitempty" db:"discord_id" json:"discord_id,omitempty"`
}

func GetUserByID[H DBStruct](h H, id string) (*User, error) {
	const op errors.Op = "GetUserByID"
	query := fmt.Sprintf("select id, discord_id from %s.user where id = ?", DiscordBotDatabase())
	var u User

	err := h.Get(&u, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "User Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &u, nil
}

func GetUserByDiscordID[H DBStruct](h H, did string) (*User, error) {
	const op errors.Op = "GetUserByDiscordID"
	query := fmt.Sprintf("select id, discord_id from %s.user where discord_id = ?", DiscordBotDatabase())
	var u User

	err := h.Get(&u, query, did)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "User Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &u, nil
}

type StakingUserType string

const (
	NormalUser  StakingUserType = "normal"
	ProjectUser StakingUserType = "project"
)

type StakingUser struct {
	ID         string          `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Address    string          `structs:"address,omitempty" db:"address" json:"address,omitempty"`
	Type       StakingUserType `structs:"type,omitempty" db:"type" json:"type,omitempty"`
	Name       *string         `structs:"name,omitempty" db:"name" json:"name,omitempty"`
	About      *string         `structs:"about,omitempty" db:"about" json:"about,omitempty"`
	Avatar     *string         `structs:"avatar,omitempty" db:"avatar" json:"avatar,omitempty"`
	AvatarID   *uint64         `structs:"avatar_id,omitempty" db:"avatar_id" json:"avatar_id,omitempty"`
	Banner     *string         `structs:"banner,omitempty" db:"banner" json:"banner,omitempty"`
	BannerID   *uint64         `structs:"banner_id,omitempty" db:"banner_id" json:"banner_id,omitempty"`
	Email      *string         `structs:"email,omitempty" db:"email" json:"email,omitempty"`
	DiscordID  *string         `structs:"discord_id,omitempty" db:"discord_id" json:"discord_id,omitempty"`
	TwitterID  *string         `structs:"twitter_id,omitempty" db:"twitter_id" json:"twitter_id,omitempty"`
	TelegramID *uint64         `structs:"telegram_id,omitempty" db:"telegram_id" json:"telegram_id,omitempty"`
	Domain     *string         `structs:"domain,omitempty" db:"domain" json:"domain,omitempty"`
}

var stakingUserTableKeySlices = map[TableKeySlice][]string{
	Tiny:  {"id", "address", "type", "name"},
	Small: {"id", "address", "type", "name", "about", "avatar", "avatar_id", "banner", "banner_id"},
	Full:  {"id", "address", "type", "name", "about", "avatar", "avatar_id", "banner", "banner_id", "email", "discord_id", "twitter_id", "telegram_id", "domain"},
}

func StakingUserTableKeys(st TableKeySlice) []string {
	return stakingUserTableKeySlices[st]
}

func GetStakingUser[H Handle](h H, id string, st TableKeySlice) (*StakingUser, error) {
	const op errors.Op = "GetUser"
	query := fmt.Sprintf("select %s from %s.user where id = ?", strings.Join(StakingUserTableKeys(st), ","), StakingDatabase())

	var u StakingUser
	err := h.Get(&u, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "User Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &u, nil
}

func GetStakingUserByAddress[H Handle](h H, address string, st TableKeySlice) (*StakingUser, error) {
	const op errors.Op = "GetUserByAddress"
	query := fmt.Sprintf("select %s from %s.user where address = ?", strings.Join(StakingUserTableKeys(st), ","), StakingDatabase())

	var u StakingUser
	err := h.Get(&u, query, address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "User Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &u, nil
}

func GetStakingUsers[H Handle](h H, st TableKeySlice, start uint64, limit uint64) (*[]StakingUser, error) {
	const op errors.Op = "GetUsers"
	query := fmt.Sprintf("select %s from %s.user limit ?, ?", strings.Join(StakingUserTableKeys(st), ","), StakingDatabase())

	var list []StakingUser
	err := h.Select(&list, query, start, limit)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetStakingUsersByAddresses[H Handle](h H, addresses []string, st TableKeySlice) (*[]StakingUser, error) {
	const op errors.Op = "GetUsersByAddresses"
	query := fmt.Sprintf("select %s from %s.user where address in (%s)", strings.Join(StakingUserTableKeys(st), ","), StakingDatabase(), strings.Repeat(",?", len(addresses)-1))

	var list []StakingUser
	err := h.Select(&list, query, misc.ToInterfaceSlice(addresses)...)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetStakingUsersWhere[H Handle](h H, st TableKeySlice, where string, args ...interface{}) (*[]StakingUser, error) {
	const op errors.Op = "GetUsersWhere"
	query := fmt.Sprintf("select %s from %s.user where %s", strings.Join(StakingUserTableKeys(st), ","), StakingDatabase(), where)

	var list []StakingUser
	err := h.Select(&list, query, args...)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetStakingPoolEntryUsersByPoolID[H Handle](h H, poolID string, st TableKeySlice, start uint64, limit uint64) (*[]StakingUser, error) {
	const op errors.Op = "GetPoolEntryUsersByPoolID"
	query := fmt.Sprintf("select %s from %s.user where user_id in (select user_id from %s.pool_entry where pool_id = ? limit ?, ? order by entry_date desc)", strings.Join(StakingUserTableKeys(st), ","), StakingDatabase(), StakingDatabase())

	var list []StakingUser
	err := h.Select(&list, query, poolID, start, limit)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}
