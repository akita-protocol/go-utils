package subscriptions

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/algorand/go-algorand-sdk/v2/client/v2/common/models"
	"gitlab.algotoolchain.com/keelakay/utils/nfd"
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

func PrintProperties(properties *SubscriptionProperties) {
	fmt.Println("Sub AppID:", properties.AppID)
	prettyJson, _ := json.MarshalIndent(properties, "", "  ")
	fmt.Println(string(prettyJson))
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
		state = SubscriptionProperties{
			State: map[string]string{},
		}
		valAsStr string
		key      string
	)

	processKeyAndVal := func(key string, valType uint64, intVal uint64, stringVal []byte) {
		switch valType {
		case 1: // bytes
			fmt.Println("bytes", key)
			if len(stringVal) == 32 && (key == "recipient" || key == "subscriber") {
				// 32 bytes and key name has .a [algorand address] suffix - parse accordingly - strip suffix
				valAsStr = nfd.RawPKAsAddress(stringVal).String()
			} else if len(stringVal) == 8 && !isStringPrintable(string(stringVal)) {
				// Assume it's a big-endian integer
				valAsStr = strconv.FormatUint(binary.BigEndian.Uint64(stringVal), 10)
			} else {
				valAsStr = string(stringVal)
			}
		case 2: // uint
			fmt.Println("uint", key)
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
