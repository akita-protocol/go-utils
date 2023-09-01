package watcher

type WatchSubscriptionCategory string

const (
	NFDAppIDs              WatchSubscriptionCategory = "NFDAppIDs"
	SubscriptionAppIDs     WatchSubscriptionCategory = "SubscriptionAppIDs"
	HoldingsWallets        WatchSubscriptionCategory = "HoldingsWallets"
	CreatorWallets         WatchSubscriptionCategory = "CreatorWallets"
	DiscordRelevantWallets WatchSubscriptionCategory = "DiscordRelevantWallets"
	AssetChanges           WatchSubscriptionCategory = "AssetChanges"
	FoundationWallets      WatchSubscriptionCategory = "FoundationWallets"
	IncWallets             WatchSubscriptionCategory = "IncWallets"
)

type WatchSubscriptions struct {
	NFDAppIDs              []uint64
	SubscriptionAppIDs     []uint64
	HoldingsWallets        []string
	CreatorWallets         []string
	DiscordRelevantWallets []string
	AssetChanges           []uint64
	FoundationWallets      []string
	IncWallets             []string
}

type WatchSubscriptionLabeledAddress struct {
	Address string
	Labels  []WatchSubscriptionCategory
}
