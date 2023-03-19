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
	NFD | *NFD | NFDWallet | *NFDWallet | NFDJson | *NFDJson | SessionToken | *SessionToken | Wallet | *Wallet | WalletToken | *WalletToken | User | *User | Transaction | *Transaction | Subscription | *Subscription | SubscriptionFactory | *SubscriptionFactory | SubscriptionBlacklist | *SubscriptionBlacklist | SubscriptionOffering | *SubscriptionOffering | SubscriptionPass | *SubscriptionPass | SubscriptionPayment | *SubscriptionPayment
}
