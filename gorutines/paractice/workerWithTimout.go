package paractice

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func workerWithTimeout(jobs <-chan int, wg *sync.WaitGroup, workerId int, ctx context.Context) {
	defer wg.Done()
	for job := range jobs {
		select {
		case <-time.After(200 * time.Millisecond):
			fmt.Printf("worker %d processed job %d\n", workerId, job)
		case <-ctx.Done():
			fmt.Printf("worker %d stopped due to timeout\n", workerId)
			return
		}
	}
}

func testWorkerWithTimeout() {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancel()

	jobs := make(chan int)
	workerCount := 3

	// запускаем воркеров
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go workerWithTimeout(jobs, wg, i, ctx)
	}

	// заполняю канал с работой
	go func() {
		for j := 1; j <= 10; j++ {
			jobs <- j
		}

		close(jobs)
	}()

	// подожду пока все воркеры доработают
	wg.Wait()

	fmt.Println("program is finished")
}
