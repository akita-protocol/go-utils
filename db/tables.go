package db

func getTable(obj interface{}) string {
	switch obj.(type) {
	case NFD, *NFD:
		return "nfd"
	case NFDJson, *NFDJson:
		return "nfd_json"
	case NFDWallet, *NFDWallet:
		return "nfd_wallet"
	case SessionToken, *SessionToken:
		return "session_token"
	case Transaction, *Transaction:
		return "txn"
	case User, *User:
		return "user"
	case Wallet, *Wallet:
		return "wallet"
	case WalletToken, *WalletToken:
		return "wallet_token"
	case Subscription, *Subscription:
		return "subscription"
	case SubscriptionFactory, *SubscriptionFactory:
		return "subscription_factory"
	case SubscriptionBlacklist, *SubscriptionBlacklist:
		return "subscription_blacklist"
	case SubscriptionOffering, *SubscriptionOffering:
		return "subscription_offering"
	case SubscriptionPass, *SubscriptionPass:
		return "subscription_pass"
	case SubscriptionPayment, *SubscriptionPayment:
		return "subscription_payment"
	default:
		return ""
	}
}
