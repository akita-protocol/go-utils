package amm

type Provider string

const (
	Tinyman Provider = "tinyman"
	Humble  Provider = "humble"
	Pact    Provider = "pactfi"
	Algofi  Provider = "algofi"
)

type Version string

const (
	Default Version = "default"
	V1_1    Version = "v1.1"
	V2      Version = "v2"
)

type Network string

const (
	Mainnet Network = "mainnet"
	Testnet Network = "testnet"
)

func GetValidatorAppID(p Provider, v Version, n Network) uint64 {
	var m = map[Provider]map[Version]map[Network]uint64{
		Tinyman: {
			V1_1: {
				Testnet: 62368684,
				Mainnet: 552635992,
			},
			V2: {
				Testnet: 148607000,
				Mainnet: 1002541853,
			},
		},
		Humble: {
			Default: {
				Testnet: 0,
				Mainnet: 0,
			},
		},
		Pact: {
			Default: {
				Testnet: 0,
				Mainnet: 0,
			},
		},
		Algofi: {
			Default: {
				Testnet: 0,
				Mainnet: 0,
			},
		},
	}

	return m[p][v][n]
}
