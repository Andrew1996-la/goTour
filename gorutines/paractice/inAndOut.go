package paractice

import (
	"fmt"
	"sync"
)

func workerInAndOut(jobsCh <-chan int, res chan<- string, wg *sync.WaitGroup, workerIndex int) {
	defer wg.Done()

	for n := range jobsCh {
		res <- fmt.Sprintf("worker %d записал %d", workerIndex, n*2)
	}
}

func workerInAndOutTest() {
	jobs := make(chan int)
	result := make(chan string)
	wg := &sync.WaitGroup{}

	go func() {
		for i := 1; i <= 20; i++ {
			jobs <- i
		}

		close(jobs)
	}()

	wg.Add(2)
	go workerInAndOut(jobs, result, wg, 1)
	go workerInAndOut(jobs, result, wg, 2)

	go func() {
		wg.Wait()
		close(result)
	}()

	for v := range result {
		fmt.Println(v)
	}
}
