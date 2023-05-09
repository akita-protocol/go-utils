package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

var NFDStates = []string{"available", "minting", "reserved", "forSale", "owned"}

type NFD struct {
	ID               string  `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	State            string  `structs:"state,omitempty" db:"state" json:"state,omitempty"`
	CurrentAsOfBlock uint64  `structs:"current_as_of_block,omitempty" db:"current_as_of_block" json:"current_as_of_block,omitempty"`
	AppID            uint64  `structs:"app_id,omitempty" db:"app_id" json:"app_id,omitempty"`
	AsaID            uint64  `structs:"asa_id,omitempty" db:"asa_id" json:"asa_id,omitempty"`
	Name             string  `structs:"name,omitempty" db:"name" json:"name,omitempty"`
	Version          string  `structs:"version,omitempty" db:"version" json:"version,omitempty"`
	UserName         *string `structs:"username,omitempty" db:"username" json:"username,omitempty"`
	Bio              *string `structs:"bio,omitempty" db:"bio" json:"bio,omitempty"`
	Email            *string `structs:"email,omitempty" db:"email" json:"email,omitempty"`
	Avatar           *string `structs:"avatar,omitempty" db:"avatar" json:"avatar,omitempty"`
	AvatarID         *uint64 `structs:"avatar_id,omitempty" db:"avatar_id" json:"avatar_id,omitempty"`
	Banner           *string `structs:"banner,omitempty" db:"banner" json:"banner,omitempty"`
	BannerID         *uint64 `structs:"banner_id,omitempty" db:"banner_id" json:"banner_id,omitempty"`
	DiscordID        *string `structs:"discord_id,omitempty" db:"discord_id" json:"discord_id,omitempty"`
	TwitterID        *string `structs:"twitter_id,omitempty" db:"twitter_id" json:"twitter_id,omitempty"`
	TelegramID       *uint64 `structs:"telegram_id,omitempty" db:"telegram_id" json:"telegram_id,omitempty"`
	Domain           *string `structs:"domain,omitempty" db:"domain" json:"domain,omitempty"`
	LastChangedBlock uint64  `structs:"last_changed_block" db:"last_changed_block" json:"last_changed_block,omitempty"`
	CreatedBlock     uint64  `structs:"created_block,omitempty" db:"created_block" json:"created_block,omitempty"`
	PurchasedBlock   uint64  `structs:"purchased_block,omitempty" db:"purchased_block" json:"purchased_block,omitempty"`
}

var NFDTableKeySlices = map[TableKeySlice][]string{
	Tiny:  {"id", "app_id", "name", "username", "avatar", "avatar_id", "banner", "banner_id"},
	Small: {"id", "app_id", "name", "username", "email", "avatar", "avatar_id", "banner", "banner_id", "discord_id", "twitter_id", "telegram_id", "domain"},
	Full:  {"id", "state", "current_as_of_block", "app_id", "asa_id", "name", "version", "username", "email", "avatar", "avatar_id", "banner", "banner_id", "discord_id", "twitter_id", "telegram_id", "domain", "last_changed_block", "created_block", "purchased_block"},
}

func NFDTableKeys(st TableKeySlice) []string {
	return NFDTableKeySlices[st]
}

func GetNFDIDByAppID[H DBStruct](h H, appID uint64, st TableKeySlice) (string, error) {
	const op errors.Op = "GetNFDIDByAppID"
	query := fmt.Sprintf("select %s from %s.nfd where app_id = ?", strings.Join(NFDTableKeys(st), ","), NFDDatabase())
	var id string

	err := h.Get(&id, query, appID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.E(pkg, op, errors.DatabaseResultNotFound, err, "NFD ID Not Found")
		}
		return "", errors.E(pkg, op, err)
	}
	return id, nil
}

func GetNFDsByAddress[H DBStruct](h H, address string, st TableKeySlice) (*[]NFD, error) {
	const op errors.Op = "GetNFDsByAddress"
	query := fmt.Sprintf("select %s from %s.nfd where id = (select id from nfd.nfd_wallet where address = ?)", strings.Join(NFDTableKeys(st), ","), NFDDatabase())
	var list []NFD

	err := h.Select(&list, query, address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "nfd not found")
		}
		return nil, errors.E(pkg, op, err)
	}

	if !(len(list) > 0) {
		return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, fmt.Errorf("nfd not found"))
	}

	return &list, nil
}

func GetAllNFDAppIDs[H DBStruct](h H) (*[]uint64, error) {
	const op errors.Op = "GetAllNFDAppIDs"
	query := fmt.Sprintf("select app_id from %s.nfd", NFDDatabase())
	var list []uint64

	err := h.Select(&list, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "nfd app ids not found")
		}
		return nil, errors.E(pkg, op, err)
	}

	if !(len(list) > 0) {
		return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, fmt.Errorf("nfd app ids not found"))
	}

	return &list, nil
}

func GetLatestNFDUpdateBlock[H DBStruct](h H) (uint64, error) {
	const op errors.Op = "GetLatestNFDUpdateBlock"
	query := fmt.Sprintf("select max(current_as_of_block) from %s.nfd", NFDDatabase())
	var block uint64

	err := h.Get(&block, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "NFD Block Not Found")
		}
		return 0, errors.E(pkg, op, err)
	}
	return block, nil
}
