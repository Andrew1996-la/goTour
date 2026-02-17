package gorutines

import (
	"context"
	"fmt"
)

func generator(source chan<- int, ctx context.Context) {
	for i := 1; i <= 100; i++ {
		select {
		case <-ctx.Done():
			close(source)
			return
		case source <- i:
		}
	}
	close(source)
}

func processor(source <-chan int, result chan<- int, ctx context.Context) {
	for v := range source {
		select {
		case <-ctx.Done():
			close(result)
			return
		case result <- v * v:
		}
	}

	close(result)
}

func aggregator(result <-chan int, sumCh chan<- int, cancel context.CancelFunc) {
	sum := 0
	for v := range result {
		if sum+v >= 200 {
			fmt.Println("aggregator result more than 200")
			cancel()
			break
		}
		sum += v
	}
	sumCh <- sum

	close(sumCh)
}

func piplineWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	source := make(chan int)
	result := make(chan int)
	sumCh := make(chan int)

	go generator(source, ctx)
	go processor(source, result, ctx)
	go aggregator(result, sumCh, cancel)

	sum := <-sumCh
	fmt.Println("sum:", sum)
	fmt.Println("finish pipline")
}
