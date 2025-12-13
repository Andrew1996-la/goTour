package tasks

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func processingJob(jobsCh chan int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case jobItem := <-jobsCh:
			time.Sleep(1 * time.Second)
			fmt.Println("it is order", jobItem)
		}
	}
}

// Задача: «Система задач с отменой»
func ProcessingWithCancelByTime() {
	context, closeContext := context.WithTimeout(context.Background(), 4*time.Second)
	wg := &sync.WaitGroup{}
	jobsCh := make(chan int)

	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	go func() {
		for _, job := range jobs {
			jobsCh <- job
		}
		close(jobsCh)
	}()

	wg.Add(1)
	go processingJob(jobsCh, context, wg)
	
	wg.Wait()
	closeContext()

	fmt.Println("Программа успешно завершена")
}
