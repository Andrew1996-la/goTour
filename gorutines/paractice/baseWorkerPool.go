package paractice

import (
	"fmt"
	"sync"
	"time"
)

func baseWorkerPool(jobs <-chan int, wg *sync.WaitGroup, workerId int) {
	defer wg.Done()
	for job := range jobs {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("worker %d processed job %d\n", workerId, job)
	}
}

func testBaseWorkerPool() {
	wg := &sync.WaitGroup{}
	jobs := make(chan int)
	workerCount := 3

	// запускаем воркеров
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go baseWorkerPool(jobs, wg, i)
	}

	// заполняю канал с работой
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	// закрою канал
	close(jobs)

	// подожду пока все воркеры доработают
	wg.Wait()

	fmt.Println("program is finished")
}
