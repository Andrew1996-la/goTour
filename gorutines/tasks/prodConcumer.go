package tasks

import (
	"fmt"
	"sync"
	"time"
)

func producers(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		time.Sleep(300 * time.Millisecond)
		ch <- i
	}
}

func consumer(ch chan int) {
	for msg := range ch {
		fmt.Println(msg)
	}
}

func ProdConsumer() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go producers(ch, wg)
	go producers(ch, wg)

	go func() {
		wg.Wait()
		close(ch)
	}()
	consumer(ch)

	fmt.Println("Программа завершена")
}
