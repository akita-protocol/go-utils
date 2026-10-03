package watcher

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func TestStartSurvivesUnknownLabelsAndPanics(t *testing.T) {
	var mu sync.Mutex
	processed := []string{}
	eq := NewEventQueue[AddressSubCategory](
		[]Processor[string]{func(item string) error {
			if item == "boom" {
				panic("processor bug")
			}
			if item == "fail" {
				return errors.New("failed")
			}
			mu.Lock()
			processed = append(processed, item)
			mu.Unlock()
			return nil
		}},
	)
	eq.Subs[HoldingsWallets] = &sync.Map{}
	go eq.Start()

	// FoundationWallets has no map: before the fix this panicked and stopped the process.
	if err := eq.AddToQueue([]WSList[AddressSubCategory, string]{
		{Labels: []AddressSubCategory{FoundationWallets, HoldingsWallets}, List: []string{"a", "boom", "fail", "b"}},
	}); err != nil {
		t.Fatal(err)
	}

	deadline := time.Now().Add(2 * time.Second)
	for {
		mu.Lock()
		n := len(processed)
		mu.Unlock()
		if n == 2 {
			break
		}
		if time.Now().After(deadline) {
			t.Fatalf("queue stopped early; processed %v", processed)
		}
		time.Sleep(5 * time.Millisecond)
	}

	if _, ok := eq.Subs[HoldingsWallets].Load("b"); !ok {
		t.Fatal("known labels must still be stored")
	}
	if !eq.isFailure("boom") || !eq.isFailure("fail") {
		t.Fatal("a panic and an error must both be recorded as failures")
	}
}
