package watcher

import (
	"fmt"
	"sync"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type Category interface {
	AddressSubCategory | AssetSubCategory | AppSubCategory
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
	FoundationWallets AddressSubCategory = "FoundationWallets"
	IncWallets        AddressSubCategory = "IncWallets"
)

type AssetSubCategory string

const (
	AssetChanges AssetSubCategory = "AssetChanges"
)

type AppSubCategory string

const (
	NFDAppIDs          AppSubCategory = "NFDAppIDs"
	SubscriptionAppIDs AppSubCategory = "SubscriptionAppIDs"
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
					eq.Subs[category].Store(item, struct{}{})
				}

				for _, process := range eq.Processors {
					err := process(item)
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
