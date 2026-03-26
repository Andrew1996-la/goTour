package anytask

import (
	"fmt"
	"sync"
)

func generator(start, end int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := start; i <= end; i++ {
			ch <- i
		}
	}()

	return ch
}

func fanIn(ch1, ch2 <-chan int) <-chan int {
	res := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch1 {
			res <- v
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch2 {
			res <- v
		}
	}()

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}

func testFanIn() {
	ch1 := generator(1, 5)
	ch2 := generator(10, 15)

	res := fanIn(ch1, ch2)

	for r := range res {
		fmt.Println(r)
	}
}
