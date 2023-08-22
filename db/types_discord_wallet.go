package db

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type DiscordWallet struct {
	ID       string  `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	UserID   string  `structs:"user_id,omitempty" db:"user_id" json:"user_id,omitempty"`
	Label    *string `structs:"label,omitempty" db:"label" json:"label,omitempty"`
	Address  string  `structs:"address,omitempty" db:"address" json:"address,omitempty"`
	Verified *bool   `structs:"verified,omitempty" db:"verified" json:"verified,omitempty"`
	Main     *bool   `structs:"main,omitempty" db:"main" json:"main,omitempty"`
}

func GetDiscordWalletByID[H DBStruct](h H, id string) (*DiscordWallet, error) {
	const op errors.Op = "GetWalletByID"
	query := fmt.Sprintf("select id, user_id, label, address, verified, main from %s.wallet where id = ?", DiscordBotDatabase())
	var w DiscordWallet

	err := h.Get(&w, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Wallet Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &w, nil
}

func GetDiscordWalletByAddress[H DBStruct](h H, id string) (*DiscordWallet, error) {
	const op errors.Op = "GetWalletByAddress"
	query := fmt.Sprintf("select id, user_id, label, address, verified, main from %s.wallet where address = ?", DiscordBotDatabase())
	var w DiscordWallet

	err := h.Get(&w, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Wallet Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &w, nil
}

func GetAllVerifiedDiscordWallets[H DBStruct](h H) (*[]DiscordWallet, error) {
	const op errors.Op = "GetAllVerifiedDiscordWallets"
	query := fmt.Sprintf("select id, user_id, label, address, verified, main from %s.wallet where verified = 1", DiscordBotDatabase())
	var list []DiscordWallet

	err := h.Select(&list, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "wallets not found")
		}
		return nil, errors.E(pkg, op, err)
	}

	if !(len(list) > 0) {
		return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, fmt.Errorf("wallets not found"))
	}

	return &list, nil
}

func GetVerifiedDiscordWalletByUserID[H DBStruct](h H, userID string) (*DiscordWallet, error) {
	const op errors.Op = "GetVerifiedWalletByUserID"
	query := fmt.Sprintf("select id, user_id, label, address, verified, main from %s.wallet where user_id = ? and verified = 1 limit 1", DiscordBotDatabase())
	var w DiscordWallet

	err := h.Get(&w, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Wallet Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &w, nil
}

func GetVerifiedDiscordWalletsByUserID[H DBStruct](h H, userID string) (*[]DiscordWallet, error) {
	const op errors.Op = "GetVerifiedWalletByUserID"
	query := fmt.Sprintf("select id, user_id, label, address, verified, main from %s.wallet where user_id = ? and verified = 1", DiscordBotDatabase())
	var list []DiscordWallet

	err := h.Select(&list, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "wallets not found")
		}
		return nil, errors.E(pkg, op, err)
	}

	if !(len(list) > 0) {
		return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, fmt.Errorf("wallets not found"))
	}

	return &list, nil
}

func DeleteDiscordWalletByAddress[H DBStruct](h H, address string) error {
	const op errors.Op = "DeleteWalletByAddress"
	query := fmt.Sprintf("delete from %s.wallet where address = ?", DiscordBotDatabase())
	var err error

	switch h := any(h).(type) {
	case *sqlx.Tx:
		stmt, err := h.Prepare(query)
		if err != nil {
			return errors.E(pkg, op, errors.Database, err, "Failed to Prepare Query")
		}

		_, err = stmt.Exec(address)
		if err != nil {
			return errors.E(pkg, op, errors.Database, err, "Failed to Execute Query")
		}
	case *sqlx.DB:
		_, err = h.Exec(query, address)
		if err != nil {
			return errors.E(pkg, op, errors.Database, err, "Failed to Execute Query")
		}
	}

	return nil
}

type DiscordWalletToken struct {
	WalletID string `structs:"wallet_id,omitempty" db:"wallet_id" json:"wallet_id,omitempty"`
	Token    string `structs:"token,omitempty" db:"token" json:"token,omitempty"`
}

func GetDiscordWalletToken[H DBStruct](h H, id string) error {
	const op errors.Op = "Get"
	query := fmt.Sprintf("select wallet_id, token from %s.wallet_token where wallet_id = ?", DiscordBotDatabase())
	var r DiscordWallet

	err := h.Get(&r, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Wallet Token Not Found")
		}
		return errors.E(pkg, op, err)
	}
	return nil
}
