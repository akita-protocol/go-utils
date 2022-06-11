package db

import (
	"database/sql"

	"gitlab.keelakay.com/keelakay/utils/errors"
	"gitlab.keelakay.com/keelakay/utils/uuid"
)

type Wallet struct {
	ID       string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	UserID   string `structs:"user_id,omitempty" db:"user_id" json:"user_id,omitempty"`
	Label    string `structs:"label,omitempty" db:"label" json:"label,omitempty"`
	Address  string `structs:"address,omitempty" db:"address" json:"address,omitempty"`
	Verified *bool  `structs:"verified,omitempty" db:"verified" json:"verified,omitempty"`
	Main     *bool  `structs:"main,omitempty" db:"main" json:"main,omitempty"`
}

func NewWallet() *Wallet {
	return &Wallet{
		ID: uuid.New(uuid.Wallet),
	}
}

func GetWalletByID[H DBStruct](h H, id string) (*Wallet, error) {
	const op errors.Op = "Get"

	const query = " select id, user_id, label, address, verified, main from wallet where id = ?"
	var w Wallet

	err := h.Get(&w, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Wallet Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &w, nil
}

func GetWalletByUserID[H DBStruct](h H, userID string) (*Wallet, error) {
	const op errors.Op = "Get"

	const query = " select id, user_id, label, address, verified, main from wallet where user_id = ?"
	var w Wallet

	err := h.Get(&w, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Wallet Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &w, nil
}

type WalletToken struct {
	WalletID string `structs:"wallet_id,omitempty" db:"wallet_id" json:"wallet_id,omitempty"`
	Token    string `structs:"token,omitempty" db:"token" json:"token,omitempty"`
}

func NewWalletToken(walletID string) *WalletToken {
	return &WalletToken{
		WalletID: walletID,
		Token:    uuid.New(uuid.Token),
	}
}

func GetWalletToken[H DBStruct](h H, id string) error {
	const op errors.Op = "Get"

	const query = " select wallet_id, token from wallet_token where wallet_id = ?"
	var r Wallet

	err := h.Get(&r, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Wallet Token Not Found")
		}
		return errors.E(pkg, op, err)
	}
	return nil
}
