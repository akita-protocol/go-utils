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
	NFD | *NFD | NFDWallet | *NFDWallet | NFDJson | *NFDJson | Collection | *Collection | CollectionJson | *CollectionJson | CollectionPrefix | *CollectionPrefix | CollectionArtist | *CollectionArtist | CollectionAsset | *CollectionAsset | CollectionExcludedAsset | *CollectionExcludedAsset | CollectionExtras | *CollectionExtras | Property | *Property | PropertyValue | *PropertyValue | PropertyValueExtras | *PropertyValueExtras | SessionToken | *SessionToken | Wallet | *Wallet | WalletToken | *WalletToken | User | *User | Transaction | *Transaction | Subscription | *Subscription | SubscriptionFactory | *SubscriptionFactory | SubscriptionBlacklist | *SubscriptionBlacklist | SubscriptionOffering | *SubscriptionOffering | SubscriptionPass | *SubscriptionPass | SubscriptionPayment | *SubscriptionPayment
}
