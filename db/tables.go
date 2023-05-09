package db

import "fmt"

func getTable(obj interface{}) string {
	switch obj.(type) {
	case NFD, *NFD:
		return fmt.Sprintf("%s.nfd", NFDDatabase())
	case NFDJson, *NFDJson:
		return fmt.Sprintf("%s.nfd_json", NFDDatabase())
	case NFDWallet, *NFDWallet:
		return fmt.Sprintf("%s.nfd_wallet", NFDDatabase())
	case Collection, *Collection:
		return fmt.Sprintf("%s.collection", NFDDatabase())
	case CollectionJson, *CollectionJson:
		return fmt.Sprintf("%s.collection_json", NFDDatabase())
	case CollectionPrefix, *CollectionPrefix:
		return fmt.Sprintf("%s.collection_prefix", NFDDatabase())
	case CollectionAsset, *CollectionAsset:
		return fmt.Sprintf("%s.collection_asset", NFDDatabase())
	case CollectionExcludedAsset, *CollectionExcludedAsset:
		return fmt.Sprintf("%s.collection_excluded_asset", NFDDatabase())
	case CollectionArtist, *CollectionArtist:
		return fmt.Sprintf("%s.collection_artist", NFDDatabase())
	case CollectionExtras, *CollectionExtras:
		return fmt.Sprintf("%s.collection_extras", NFDDatabase())
	case SessionToken, *SessionToken:
		return fmt.Sprintf("%s.session_token", DiscordBotDatabase())
	case User, *User:
		return fmt.Sprintf("%s.user", DiscordBotDatabase())
	case Wallet, *Wallet:
		return fmt.Sprintf("%s.wallet", DiscordBotDatabase())
	case WalletToken, *WalletToken:
		return fmt.Sprintf("%s.wallet_token", DiscordBotDatabase())
	case Subscription, *Subscription:
		return fmt.Sprintf("%s.subscription", SubscriptionDatabase())
	case SubscriptionFactory, *SubscriptionFactory:
		return fmt.Sprintf("%s.subscription_factory", SubscriptionDatabase())
	case SubscriptionBlacklist, *SubscriptionBlacklist:
		return fmt.Sprintf("%s.subscription_blacklist", SubscriptionDatabase())
	case SubscriptionOffering, *SubscriptionOffering:
		return fmt.Sprintf("%s.subscription_offering", SubscriptionDatabase())
	case SubscriptionPass, *SubscriptionPass:
		return fmt.Sprintf("%s.subscription_pass", SubscriptionDatabase())
	case SubscriptionPayment, *SubscriptionPayment:
		return fmt.Sprintf("%s.subscription_payment", SubscriptionDatabase())
	default:
		return ""
	}
}
