package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"gitlab.keelakay.com/keelakay/utils/errors"
)

type Wallet struct {
	ID       string  `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	UserID   string  `structs:"user_id,omitempty" db:"user_id" json:"user_id,omitempty"`
	Label    *string `structs:"label,omitempty" db:"label" json:"label,omitempty"`
	Address  string  `structs:"address,omitempty" db:"address" json:"address,omitempty"`
	Verified *bool   `structs:"verified,omitempty" db:"verified" json:"verified,omitempty"`
	Main     *bool   `structs:"main,omitempty" db:"main" json:"main,omitempty"`
}

func GetWalletByID[H DBStruct](h H, id string) (*Wallet, error) {
	const op errors.Op = "GetWalletByID"
	const query = "select id, user_id, label, address, verified, main from wallet where id = ?"
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

func GetWalletByAddress[H DBStruct](h H, id string) (*Wallet, error) {
	const op errors.Op = "GetWalletByAddress"
	const query = "select id, user_id, label, address, verified, main from wallet where address = ?"
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

func GetVerifiedWalletByUserID[H DBStruct](h H, userID string) (*Wallet, error) {
	const op errors.Op = "GetVerifiedWalletByUserID"
	const query = "select id, user_id, label, address, verified, main from wallet where user_id = ? and verified = 1 limit 1"
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

func GetVerifiedWalletsByUserID[H DBStruct](h H, userID string) (*[]Wallet, error) {
	const op errors.Op = "GetVerifiedWalletByUserID"
	const query = "select id, user_id, label, address, verified, main from wallet where user_id = ? and verified = 1"
	var list []Wallet

	err := h.Select(&list, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Wallet Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func DeleteWalletByAddress[H DBStruct](h H, address string) error {
	const op errors.Op = "DeleteWalletByAddress"
	const query = "delete from wallet where address = ?"
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

type WalletToken struct {
	WalletID string `structs:"wallet_id,omitempty" db:"wallet_id" json:"wallet_id,omitempty"`
	Token    string `structs:"token,omitempty" db:"token" json:"token,omitempty"`
}

func GetWalletToken[H DBStruct](h H, id string) error {
	const op errors.Op = "Get"

	const query = "select wallet_id, token from wallet_token where wallet_id = ?"
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
