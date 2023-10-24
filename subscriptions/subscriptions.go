package subscriptions

func GetFactoryAppID(isProduction bool) uint64 {
	if isProduction {
		return 166131287
	}
	return 166131287
}
