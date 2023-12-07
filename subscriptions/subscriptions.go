package subscriptions

import (
	"encoding/base64"
	"encoding/binary"
	"strconv"

	"github.com/algorand/go-algorand-sdk/v2/client/v2/common/models"
)

func GetFactoryAppID(isProduction bool) uint64 {
	if isProduction {
		return 166131287
	}
	return 166131287
}

type SubscriptionProperties struct {
	AppID uint64            `json:"appID"`
	State map[string]string `json:"state"`
}

func FetchAllStateAsSubscriptionProperties(appState []models.TealKeyValue) SubscriptionProperties {
	isStringPrintable := func(str string) bool {
		for _, strRune := range str {
			if !strconv.IsPrint(strRune) {
				return false
			}
		}
		return true
	}
	var (
		state    SubscriptionProperties
		valAsStr string
		key      string
	)

	processKeyAndVal := func(key string, valType uint64, intVal uint64, stringVal []byte) {
		switch valType {
		case 1: // bytes
			if len(stringVal) == 8 && !isStringPrintable(string(stringVal)) {
				// Assume it's a big-endian integer
				valAsStr = strconv.FormatUint(binary.BigEndian.Uint64(stringVal), 10)
			} else {
				valAsStr = string(stringVal)
			}
		case 2: // uint
			valAsStr = strconv.FormatUint(intVal, 10)
		default:
			valAsStr = "unknown"
		}
		state.State[key] = valAsStr
	}

	for _, kv := range appState {
		rawKey, _ := base64.StdEncoding.DecodeString(kv.Key)
		key = string(rawKey)
		if kv.Value.Type == 1 {
			value, _ := base64.StdEncoding.DecodeString(kv.Value.Bytes)
			processKeyAndVal(key, kv.Value.Type, kv.Value.Uint, value)
		} else {
			processKeyAndVal(key, kv.Value.Type, kv.Value.Uint, nil)
		}
	}

	return state
}
