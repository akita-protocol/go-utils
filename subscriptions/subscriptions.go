package subscriptions

func GetFactoryAppID(env string) uint64 {
	if env == "staging" {
		return 166131287
	}
	return 166131287
}
