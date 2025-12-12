package tasks

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func pinger(ch chan string, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			ch <- "Ping"
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func ponger(ch chan string, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			ch <- "Pong"
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// “Пинг в канале”
func PingPong() {
	context, cancelContext := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	pingCh := make(chan string)
	pongCh := make(chan string)

	wg.Add(1)
	go pinger(pingCh, context, wg)

	wg.Add(1)
	go ponger(pongCh, context, wg)

	go func() {
		time.Sleep(5 * time.Second)
		cancelContext()
	}()

	readLoop:
	for {
		select {
		case <-context.Done():
			break readLoop
		case ping := <-pingCh:
			fmt.Println(ping)
		case pong := <-pongCh:
			fmt.Println(pong)
		}
	}

	wg.Wait()
	fmt.Println("Программа завершилась")
}
