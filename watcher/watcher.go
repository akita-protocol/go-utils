package watcher

import (
	"fmt"
	"sync"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type Category interface {
	AddressSubCategory | AssetSubCategory | AppSubCategory | TxnSubCategory
}

type ListType interface {
	string | uint64
}

// Watch Subscription Catagories are used to determine what type of data is being sent to over our messaging queue
type AddressSubCategory string

const (
	HoldingsWallets   AddressSubCategory = "HoldingsWallets"
	CreatorWallets    AddressSubCategory = "CreatorWallets"
	DiscordWallets    AddressSubCategory = "DiscordWallets"
	LpWallets         AddressSubCategory = "LpWallets"
	FoundationWallets AddressSubCategory = "FoundationWallets"
	IncWallets        AddressSubCategory = "IncWallets"
)

type AssetSubCategory string

const (
	AssetChanges AssetSubCategory = "AssetChanges"
	LPAssets     AssetSubCategory = "LpAssets"
)

type AppSubCategory string

const (
	NFDAppIDs          AppSubCategory = "NFDAppIDs"
	SubscriptionAppIDs AppSubCategory = "SubscriptionAppIDs"
)

type TxnSubCategory string

const (
	ConfirmTxn TxnSubCategory = "confirmTxn"
)

type WSList[C Category, L ListType] struct {
	Labels []C
	List   []L
}

func NewEventQueue[C Category, L ListType](processors []Processor[L]) *EventQueue[C, L] {
	return &EventQueue[C, L]{
		Subs:       map[C]*sync.Map{},
		Processors: processors,
		ListChan:   make(chan []WSList[C, L], 1000),
		Failures: QueueFailures[L]{
			Mutex: &sync.Mutex{},
			List:  map[L]error{},
		},
	}
}

type QueueFailures[L ListType] struct {
	*sync.Mutex
	List map[L]error
}

type Processor[L ListType] func(item L) error

type EventQueue[C Category, L ListType] struct {
	Subs       map[C]*sync.Map
	Processors []Processor[L]
	ListChan   chan []WSList[C, L]
	Failures   QueueFailures[L]
}

func (eq *EventQueue[C, L]) isFailure(lt L) bool {
	eq.Failures.Lock()
	defer eq.Failures.Unlock()
	_, exists := eq.Failures.List[lt]
	return exists
}

func (eq *EventQueue[C, L]) LogFailure(lt L, err error) {
	eq.Failures.Lock()
	eq.Failures.List[lt] = err
	eq.Failures.Unlock()
}

func (eq *EventQueue[C, L]) AddToQueue(list []WSList[C, L]) error {
	const op errors.Op = "AddToQueue"
	select {
	case eq.ListChan <- list:
		return nil
	default:
		return errors.E(op, fmt.Errorf("failed to add list to queue channel"))
	}
}

func (eq *EventQueue[C, L]) Start() {
	for list := range eq.ListChan {
		for _, entry := range list {
			for _, item := range entry.List {
				if eq.isFailure(item) {
					continue
				}

				for _, category := range entry.Labels {
					// Labels can come from other services over NATS. A label with
					// no map used to panic here and stop the whole process.
					sub := eq.Subs[category]
					if sub == nil {
						fmt.Printf("[EventQueue] ignoring unknown label %v for %v\n", category, item)
						continue
					}
					sub.Store(item, struct{}{})
				}

				for _, process := range eq.Processors {
					err := runProcessor(process, item)
					if err != nil {
						fmt.Println(err)
						eq.LogFailure(item, err)
						continue
					}
				}
			}
		}
	}
}

// runProcessor turns a panic in a processor into an error. Start runs in its own
// goroutine, so an unrecovered panic would stop the whole process.
func runProcessor[L ListType](process Processor[L], item L) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("event queue processor panicked on %v: %v", item, r)
		}
	}()
	return process(item)
}
