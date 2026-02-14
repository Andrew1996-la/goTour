package paractice

import (
	"context"
	"fmt"
)

func worker(ch chan<- int, ctx context.Context) {
	iterator := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker done")
			close(ch)
			return
		case ch <- iterator:
			iterator++
		}
	}
}

func testWorker() {
	ctx, cancelCtx := context.WithCancel(context.Background())
	countCh := make(chan int)

	go worker(countCh, ctx)

	for count := range countCh {
		if count == 5 {
			cancelCtx()
		}
		fmt.Println("count:", count)
	}

}
