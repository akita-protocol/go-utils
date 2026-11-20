package db

import (
	"database/sql"
	"fmt"

	"github.com/akita-protocol/go-utils/errors"
	"github.com/jmoiron/sqlx"
)

type FoundationWatchWalletType string

const (
	FoundationWatchWalletTypeFoundation FoundationWatchWalletType = "foundation"
	FoundationWatchWalletTypeInc        FoundationWatchWalletType = "inc"
	FoundationWatchWalletTypeCEX        FoundationWatchWalletType = "cex"
)

type FoundationWatchWallet struct {
	Address string                    `json:"address"`
	Type    FoundationWatchWalletType `json:"type"`
}

func GetAllFoundationWallets[H DBStruct](h H) ([]string, error) {
	const op errors.Op = "GetAllFoundationWallets"
	query := fmt.Sprintf("select address from %v.wallet where type = 'foundation'", FoundationWatchDatabase())
	var addresses []string

	err := h.Select(&addresses, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return addresses, nil
		}
		return addresses, errors.E(pkg, op, err)
	}
	return addresses, nil
}

func GetAllIncWallets[H DBStruct](h H) ([]string, error) {
	const op errors.Op = "GetAllIncWallets"
	query := fmt.Sprintf("select address from %v.wallet where type = 'inc'", FoundationWatchDatabase())
	var addresses []string

	err := h.Select(&addresses, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return addresses, nil
		}
		return addresses, errors.E(pkg, op, err)
	}
	return addresses, nil
}

func GetAllCEXWallets[H DBStruct](h H) ([]string, error) {
	const op errors.Op = "GetAllCEXWallets"
	query := fmt.Sprintf("select address from %v.wallet where type = 'cex'", FoundationWatchDatabase())
	var addresses []string

	err := h.Select(&addresses, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return addresses, nil
		}
		return addresses, errors.E(pkg, op, err)
	}
	return addresses, nil
}

func AddWallet[H DBStruct](h H, address string, t FoundationWatchWalletType) error {
	const op errors.Op = "AddFoundationWallet"
	query := fmt.Sprintf("insert into %v.wallet (address, type) values (?, ?)", FoundationWatchDatabase())
	var err error

	switch h := any(h).(type) {
	case *sqlx.Tx:
		stmt, err := h.Prepare(query)
		if err != nil {
			return errors.E(pkg, op, errors.Database, err, "Failed to Prepare Query")
		}

		_, err = stmt.Exec(address, t)
		if err != nil {
			return errors.E(pkg, op, errors.Database, err, "Failed to Execute Query")
		}
	case *sqlx.DB:
		_, err = h.Exec(query, address, t)
		if err != nil {
			return errors.E(pkg, op, errors.Database, err, "Failed to Execute Query")
		}
	}

	return nil
}
