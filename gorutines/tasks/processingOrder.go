package tasks

import (
	"fmt"
	"sync"
	"time"
)

func processingOrder(workerId int, jobCh chan int, resCh chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for v := range jobCh {
		res := fmt.Sprintf("Worker %dобработал заказа номер %d\n", workerId, v)
		resCh <- res
		time.Sleep(2 * time.Second)
	}
}

// Задача: «Параллельная обработка заказов»
func ProcessingOrder() {
	jobCh := make(chan int)
	resultCh := make(chan string, 10)

	wg := &sync.WaitGroup{}

	orders := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	go func() {
		for numberOrder := range orders {
			jobCh <- numberOrder
		}
		close(jobCh)
	}()

	wg.Add(3)
	go processingOrder(1, jobCh, resultCh, wg)
	go processingOrder(2, jobCh, resultCh, wg)
	go processingOrder(3, jobCh, resultCh, wg)

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for orderRes := range resultCh {
		fmt.Println(orderRes)
	}

	fmt.Printf("Программа завершена")
}
