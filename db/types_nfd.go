package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
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
	// ID is the same as NFD id
	ID       string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	WalletID string `structs:"wallet_id,omitempty" db:"wallet_id" json:"wallet_id,omitempty"`
}

func GetNFDWallets[H DBStruct](h H, nfdID string) (*[]Wallet, error) {
	const op errors.Op = "GetWalletByID"
	const query = "select id, user_id, label, address, verified, main from wallet where id in (select wallet_id from nfd_wallet where nfd_id = ?)"
	var list []Wallet

	err := h.Select(&list, query, nfdID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Wallets Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func DeleteNFDWalletByWalletID[H DBStruct](h H, walletID string) error {
	const op errors.Op = "DeleteWalletByAddress"
	const query = "delete from nfd_wallet where wallet_id = ?"
	var err error

	switch h := any(h).(type) {
	case *sqlx.Tx:
		stmt, err := h.Prepare(query)
		if err != nil {
			return errors.E(pkg, op, errors.Database, err, "Failed to Prepare Query")
		}

		_, err = stmt.Exec(walletID)
		if err != nil {
			return errors.E(pkg, op, errors.Database, err, "Failed to Execute Query")
		}
	case *sqlx.DB:
		_, err = h.Exec(query, walletID)
		if err != nil {
			return errors.E(pkg, op, errors.Database, err, "Failed to Execute Query")
		}
	}

	return nil
}

type NFDJson struct {
	NFDID string `structs:"nfd_id,omitempty" db:"nfd_id" json:"nfd_id,omitempty"`
	Data  string `structs:"data,omitempty" db:"data" json:"data,omitempty"`
}
