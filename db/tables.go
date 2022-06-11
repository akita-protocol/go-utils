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
	case UserNFD, *UserNFD:
		return "user_nfd"
	case Wallet, *Wallet:
		return "wallet"
	case WalletToken, *WalletToken:
		return "wallet_token"
	default:
		return ""
	}
}
