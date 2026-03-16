package paractice

import (
	"fmt"
	"sync"
	"time"
)

func fanInOutWorker(jobs <-chan int, res chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		time.Sleep(200 * time.Millisecond)
		res <- job * 2
	}
}

func fanInOutWorkerTest() {
	jobs := make(chan int)
	results := make(chan int)

	wg := &sync.WaitGroup{}

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go fanInOutWorker(jobs, results, wg)
	}

	go func() {
		for i := 1; i <= 100; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println(r)
	}

	fmt.Println("finish")
}
