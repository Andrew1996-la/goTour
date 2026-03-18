package paractice

import (
	"fmt"
	"sync"
	"time"
)

func baseSemaphore() {
	sem := make(chan struct{}, 3)
	wg := &sync.WaitGroup{}

	for i := 1; i <= 10; i++ {
		sem <- struct{}{} // Занимаю слот горутины

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defer func() {<-sem}()

			fmt.Println("start", i)
			time.Sleep(1 * time.Second)
			fmt.Println("finish", i)

		}(i)
	}

	// time.Sleep(5 * time.Second)
	wg.Wait()

	fmt.Println("program finished")
}
