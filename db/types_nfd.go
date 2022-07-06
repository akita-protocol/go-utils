package db

import (
	"database/sql"

	"gitlab.keelakay.com/keelakay/utils/errors"
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

func GetNFDIDByAppID[H DBStruct](h H, appID uint64) (string, error) {
	const op errors.Op = "GetNFDIDByAppID"
	const query = "select id from nfd where app_id = ?"
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

type NFDWallet struct {
	NFDID    string `structs:"nfd_id,omitempty" db:"nfd_id" json:"nfd_id,omitempty"`
	WalletID string `structs:"wallet_id,omitempty" db:"wallet_id" json:"wallet_id,omitempty"`
}

type NFDJson struct {
	NFDID string `structs:"nfd_id,omitempty" db:"nfd_id" json:"nfd_id,omitempty"`
	Data  string `structs:"data,omitempty" db:"data" json:"data,omitempty"`
}
