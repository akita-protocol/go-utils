package db

import (
	"github.com/jmoiron/sqlx"
)

type DBStruct interface {
	*sqlx.DB | *sqlx.Tx
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

type DBObject interface {
	NFD | NFDWallet | NFDJson | SessionToken | Wallet | WalletToken | User | UserNFD | Transaction
}
