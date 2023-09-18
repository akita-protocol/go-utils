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
	*Asset | *LP | *AssetProperty | *NFD | *NFDWallet | *NFDJson | *Community | *CommunityJson | *CommunityAssociate | *CommunityToken | *CommunityTokenDistribution | *CommunityTokenDistributionCollection | *CommunityFaq | *CommunityExtras | *Collection | *CollectionPrefix | *CollectionArtist | *CollectionAsset | *CollectionExcludedAsset | *CollectionExtras | *Property | *PropertyValue | *PropertyValueExtras | *DiscordSessionToken | *DiscordWallet | *DiscordWalletToken | *DiscordUser | *DiscordServer | *DiscordServerRole | *DiscordRoleReq | *DiscordTransaction | *Subscription | *SubscriptionFactory | *SubscriptionBlacklist | *SubscriptionOffering | *SubscriptionPass | *SubscriptionPayment | *StakingUser | *StakingPool | *StakingPoolAsset | *StakingExclusion | *StakingEntryRequirement | *StakingPoolEntry | *StakingEntryAsset | *StakingSnapshot | *StakingSnapshotAsset | *Reward | *Disbursement | *RewardCode | *StakingMerchant | *Shuffle | *ShuffleWhitelistToken | *ShuffleAsset | *ShuffleReserveAsset | *ShuffleSale | *ShufflePaymentOption | *ShuffleAssetProperty | *Holdings
}
