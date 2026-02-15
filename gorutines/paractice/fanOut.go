package paractice

import (
	"fmt"
	"sync"
)

// fanOut это когда одни канал, но много читателей
func fanOutWorker(jobs <-chan int, workerId int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range jobs {
		fmt.Printf("worker %d got %d\n", workerId, n)
	}
}

func testFanOut() {
	jobCh := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go fanOutWorker(jobCh, i, wg)
	}

	for i := 1; i <= 5; i++ {
		jobCh <- i
	}
	close(jobCh)
	wg.Wait()

	fmt.Println("Fan out is done")
}
