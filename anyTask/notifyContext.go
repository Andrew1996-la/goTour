package anytask

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func processed(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for x := range jobs {
		fmt.Printf("ID: %d started %d\n", id, x)
		time.Sleep(4 * time.Second)
		fmt.Printf("ID: %d processed %d\n", id, x)
	}

	fmt.Printf("worker %d: done\n", id)
}

func testProcessed() {
	wg := &sync.WaitGroup{}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	jobs := make(chan int)


	go func() {
		for i := 1; i <= 100; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Stop producing jobs")
				close(jobs)
				return
			case jobs <- i:
			}
		}
		close(jobs)
	}()

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go processed(i, jobs, wg)
	}

	wg.Wait()

	fmt.Println("Program finished")

}
