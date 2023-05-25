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
	case Community, *Community:
		return fmt.Sprintf("%s.community", NFDDatabase())
	case CommunityJson, *CommunityJson:
		return fmt.Sprintf("%s.community_json", NFDDatabase())
	case CommunityToken, *CommunityToken:
		return fmt.Sprintf("%s.community_token", NFDDatabase())
	case CommunityTokenDistribution, *CommunityTokenDistribution:
		return fmt.Sprintf("%s.community_token_distribution", NFDDatabase())
	case CommunityTokenDistributionCollection, *CommunityTokenDistributionCollection:
		return fmt.Sprintf("%s.community_token_distribution_collection", NFDDatabase())
	case CommunityFaq, *CommunityFaq:
		return fmt.Sprintf("%s.community_faq", NFDDatabase())
	case CommunityExtras, *CommunityExtras:
		return fmt.Sprintf("%s.community_extras", NFDDatabase())
	case Collection, *Collection:
		return fmt.Sprintf("%s.collection", NFDDatabase())
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
	case Property, *Property:
		return fmt.Sprintf("%s.property", NFDDatabase())
	case PropertyValue, *PropertyValue:
		return fmt.Sprintf("%s.property_value", NFDDatabase())
	case PropertyValueExtras, *PropertyValueExtras:
		return fmt.Sprintf("%s.property_value_extras", NFDDatabase())
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
