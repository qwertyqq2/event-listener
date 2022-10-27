package main

import (
	"sync"
	"time"

	"github.com/qwertyqq2/event-listener/internal/event-listener/provider"

	eventrun "github.com/qwertyqq2/event-listener/internal/event-run"
)

var (
	i  int64
	wg sync.WaitGroup
)

func main() {
	wg.Add(1)
	u := eventrun.U
	i = 10
	go func(idx *int64) {
		for {
			u.Deposit(*idx % 5)
			time.Sleep(1 * time.Second)
			*idx += 1
		}
	}(&i)
	time.Sleep(time.Second)

	provider := &provider.Provider{}
	provider.SetUp()
	go provider.ListenToEvent()

	wg.Wait()
}
