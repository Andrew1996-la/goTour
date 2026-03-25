package paractice

import (
	"fmt"
	"sync"
)

func simpleWorker(jobs <-chan int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		result <- job * job
	}
}

func simpleWorkerTest() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	workerCount := 3
	wg := &sync.WaitGroup{}

	jobs := make(chan int)
	result := make(chan int, len(jobs))

	go func() {
		for _, v := range nums {
			jobs <- v
		}
		close(jobs)
	}()

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go simpleWorker(jobs, result, wg)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	for res := range result {
		fmt.Println(res)
	}
	
}
