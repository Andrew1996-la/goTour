package paractice

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func workerWithTimeoutAndStat(
	jobs <-chan int,
	wg *sync.WaitGroup,
	workerId int,
	ctx context.Context,
	mutex *sync.Mutex,
	stats map[int]int,
) {
	defer wg.Done()
	// счетчик выполненной работы
	processed := 0

	for job := range jobs {
		select {
		case <-time.After(300 * time.Millisecond):
			fmt.Printf("worker %d processed job %d\n", workerId, job)
			processed++
		case <-ctx.Done():
			fmt.Printf("worker %d stopped due to timeout\n", workerId)
			mutex.Lock()
			stats[workerId] = processed
			mutex.Unlock()
			return
		}
	}

	mutex.Lock()
	stats[workerId] = processed
	mutex.Unlock()
}

func testWorkerWithTimeoutAndStat() {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	stats := make(map[int]int)

	jobs := make(chan int, 5)
	workerCount := 3

	// запускаем воркеров
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go workerWithTimeoutAndStat(jobs, wg, i, ctx, mu, stats)
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

	// отображение статы
	fmt.Println("========== STATS ==========")
	for k, v := range stats {
		fmt.Printf("worker %d processed job %d\n", k, v)
	}

	fmt.Println("program is finished")
}
