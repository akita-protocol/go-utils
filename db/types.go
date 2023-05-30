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
	Asset | *Asset | NFD | *NFD | NFDWallet | *NFDWallet | NFDJson | *NFDJson | Community | *Community | CommunityJson | *CommunityJson | *CommunityToken | CommunityToken | *CommunityTokenDistribution | CommunityTokenDistribution | *CommunityTokenDistributionCollection | CommunityTokenDistributionCollection | CommunityFaq | *CommunityFaq | CommunityExtras | *CommunityExtras | Collection | *Collection | CollectionPrefix | *CollectionPrefix | CollectionArtist | *CollectionArtist | CollectionAsset | *CollectionAsset | CollectionExcludedAsset | *CollectionExcludedAsset | CollectionExtras | *CollectionExtras | Property | *Property | PropertyValue | *PropertyValue | PropertyValueExtras | *PropertyValueExtras | SessionToken | *SessionToken | Wallet | *Wallet | WalletToken | *WalletToken | User | *User | Transaction | *Transaction | Subscription | *Subscription | SubscriptionFactory | *SubscriptionFactory | SubscriptionBlacklist | *SubscriptionBlacklist | SubscriptionOffering | *SubscriptionOffering | SubscriptionPass | *SubscriptionPass | SubscriptionPayment | *SubscriptionPayment
}
