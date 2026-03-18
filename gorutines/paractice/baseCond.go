package paractice

import (
	"fmt"
	"sync"
	"time"
)

func baseCond() {
	ready := false
	mu := &sync.Mutex{}
	cond := sync.NewCond(mu)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()

		for !ready {
			fmt.Println("Жду...")
			cond.Wait()
		}

		fmt.Println("Работа началась")
		mu.Unlock()

	}()

	time.Sleep(2 * time.Second)

	mu.Lock()
	ready = true
	cond.Signal()
	mu.Unlock()

	wg.Wait()
}
