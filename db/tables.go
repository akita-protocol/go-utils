package db

import "fmt"

func getTable(obj interface{}) string {
	switch obj.(type) {
	case Asset, *Asset:
		return fmt.Sprintf("%s.asset", NFDDatabase())
	case LP, *LP:
		return fmt.Sprintf("%s.lp", NFDDatabase())
	case AssetProperty, *AssetProperty:
		return fmt.Sprintf("%s.asset_property", NFDDatabase())
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
	case CommunityAssociate, *CommunityAssociate:
		return fmt.Sprintf("%s.community_associate", NFDDatabase())
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

	case DiscordSessionToken, *DiscordSessionToken:
		return fmt.Sprintf("%s.session_token", DiscordBotDatabase())
	case DiscordUser, *DiscordUser:
		return fmt.Sprintf("%s.user", DiscordBotDatabase())
	case DiscordWallet, *DiscordWallet:
		return fmt.Sprintf("%s.wallet", DiscordBotDatabase())
	case DiscordWalletToken, *DiscordWalletToken:
		return fmt.Sprintf("%s.wallet_token", DiscordBotDatabase())
	case DiscordServer, *DiscordServer:
		return fmt.Sprintf("%s.server", DiscordBotDatabase())
	case DiscordServerRole, *DiscordServerRole:
		return fmt.Sprintf("%s.server_role", DiscordBotDatabase())
	case DiscordRoleReq, *DiscordRoleReq:
		return fmt.Sprintf("%s.role_req", DiscordBotDatabase())
	case DiscordTransaction, *DiscordTransaction:
		return fmt.Sprintf("%s.transaction", DiscordBotDatabase())

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

	case Disbursement, *Disbursement:
		return fmt.Sprintf("%s.disbursement", StakingDatabase())
	case StakingEntryAsset, *StakingEntryAsset:
		return fmt.Sprintf("%s.entry_asset", StakingDatabase())
	case StakingEntryRequirement, *StakingEntryRequirement:
		return fmt.Sprintf("%s.entry_requirement", StakingDatabase())
	case StakingEntryRequirementAsset, *StakingEntryRequirementAsset:
		return fmt.Sprintf("%s.entry_requirement_asset", StakingDatabase())
	case StakingEntryRequirementSelection, *StakingEntryRequirementSelection:
		return fmt.Sprintf("%s.entry_requirement_selection", StakingDatabase())
	case StakingExclusion, *StakingExclusion:
		return fmt.Sprintf("%s.exclusion", StakingDatabase())
	case StakingMerchant, *StakingMerchant:
		return fmt.Sprintf("%s.merchant", StakingDatabase())
	case StakingPool, *StakingPool:
		return fmt.Sprintf("%s.pool", StakingDatabase())
	case StakingPoolAsset, *StakingPoolAsset:
		return fmt.Sprintf("%s.pool_asset", StakingDatabase())
	case StakingPoolEntry, *StakingPoolEntry:
		return fmt.Sprintf("%s.pool_entry", StakingDatabase())
	case Reward, *Reward:
		return fmt.Sprintf("%s.reward", StakingDatabase())
	case RewardCode, *RewardCode:
		return fmt.Sprintf("%s.reward_code", StakingDatabase())
	case StakingSnapshot, *StakingSnapshot:
		return fmt.Sprintf("%s.snapshot", StakingDatabase())
	case StakingSnapshotAsset, *StakingSnapshotAsset:
		return fmt.Sprintf("%s.snapshot_asset", StakingDatabase())
	case StakingUser, *StakingUser:
		return fmt.Sprintf("%s.user", StakingDatabase())

	case Shuffle, *Shuffle:
		return fmt.Sprintf("%s.shuffle", ShuffleDatabase())
	case ShuffleWhitelistToken, *ShuffleWhitelistToken:
		return fmt.Sprintf("%s.shuffle_whitelist_token", ShuffleDatabase())
	case ShuffleAsset, *ShuffleAsset:
		return fmt.Sprintf("%s.shuffle_asset", ShuffleDatabase())
	case ShuffleReserveAsset, *ShuffleReserveAsset:
		return fmt.Sprintf("%s.shuffle_reserve_asset", ShuffleDatabase())
	case ShuffleSale, *ShuffleSale:
		return fmt.Sprintf("%s.shuffle_sale", ShuffleDatabase())
	case ShufflePaymentOption, *ShufflePaymentOption:
		return fmt.Sprintf("%s.shuffle_payment_option", ShuffleDatabase())
	case ShuffleAssetProperty, *ShuffleAssetProperty:
		return fmt.Sprintf("%s.shuffle_asset_property", ShuffleDatabase())

	case Holdings, *Holdings:
		return fmt.Sprintf("%s.holdings", HoldingsDatabase())
	default:
		return ""
	}
}
