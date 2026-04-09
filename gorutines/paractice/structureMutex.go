package paractice

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu      sync.Mutex
	counter int
}

func (c *Counter) Increase() {
	c.mu.Lock()
	c.counter++
	c.mu.Unlock()
}

func structureMutexTest() {
	wg := &sync.WaitGroup{}
	var counter1, counter2 Counter

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				counter1.Increase()
				counter2.Increase()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(counter1.counter, counter2.counter)

}
