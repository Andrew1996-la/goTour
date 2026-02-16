package paractice

import (
	"context"
	"fmt"
	"sync"
)

func workerWithContext(ctx context.Context, wg *sync.WaitGroup, in <-chan int, res chan<- int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case n, ok := <-in:
			if !ok {
				return
			}
			res <- n * 2
		}
	}
}

func workerWithContextTest() {
	wg := &sync.WaitGroup{}
	ctx, cancelCtx := context.WithCancel(context.Background())

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	chList := []chan int{ch1, ch2, ch3}

	resCh := make(chan int)

	// запускаю воркеры
	for _, ch := range chList {
		wg.Add(1)
		go workerWithContext(ctx, wg, ch, resCh)
	}

	// записываю в каждый канал по 100 значений
	go func() {
		for i := 1; i <= 100; i++ {
			ch1 <- i
			ch2 <- i
			ch3 <- i
		}

		close(ch1)
		close(ch2)
		close(ch3)
	}()

	// ожидаю окончания работу горутин и закрываю канал результатов для безопасного чтения из него
	go func() {
		wg.Wait()
		close(resCh)
	}()

	resCount := 0
	for v := range resCh {
		fmt.Println("value resCh:", v)
		resCount++

		if resCount >= 10 {
			cancelCtx()
		}
	}

	fmt.Println("is done")
}
