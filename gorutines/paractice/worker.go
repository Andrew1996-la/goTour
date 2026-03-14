package paractice

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("worker %d processed job %d\n", id, job)
	}
}

func workerTest() {
	jobs := make(chan int)
	wg := &sync.WaitGroup{}

	workerCount := 3
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, jobs, wg)
	}

	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()

	fmt.Println("all jobs processed")
}
