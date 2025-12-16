package tasks

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func longRunningTask(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		return
	case <-time.After(7 * time.Second):
		fmt.Println("Работа завершена")
	}
}

func TimeoutStopGorutine() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	wg := &sync.WaitGroup{}
	defer cancel()

	wg.Add(1)
	go longRunningTask(ctx, wg)

	wg.Wait()
	fmt.Println("Завершение программы")
}
