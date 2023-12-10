package subscriptions

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/algorand/go-algorand-sdk/v2/client/v2/common/models"
	"gitlab.algotoolchain.com/keelakay/utils/db"
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

func (p *SubscriptionProperties) ToDBSubscription(round uint64) db.Subscription {
	startDateStr := p.State["start_date"]
	startDateInt, _ := strconv.ParseInt(startDateStr, 10, 64)
	startDate := time.Unix(startDateInt, 0)

	boxIndexStr := p.State["box_index"]
	boxIndex, _ := strconv.ParseUint(boxIndexStr, 10, 64)

	intervalStr := p.State["interval"]
	interval, _ := strconv.ParseUint(intervalStr, 10, 64)

	lastPaymentStr := p.State["last_payment"]
	lastPaymentInt, _ := strconv.ParseInt(lastPaymentStr, 10, 64)
	lastPayment := time.Unix(lastPaymentInt, 0)

	asaID, _ := strconv.ParseUint(p.State["asa"], 10, 64)

	amt, _ := strconv.ParseUint(p.State["amount"], 10, 64)

	passCount, _ := strconv.ParseUint(p.State["pass_count"], 10, 64)

	return db.Subscription{
		ID:               p.AppID,
		StartDate:        startDate,
		BoxIndex:         &boxIndex,
		Interval:         interval,
		LastPayment:      lastPayment,
		Subscriber:       p.State["subscriber"],
		Recipient:        p.State["recipient"],
		ASA:              &asaID,
		Amount:           amt,
		PassCount:        &passCount,
		CurrentAsOfBlock: round,
	}
}
