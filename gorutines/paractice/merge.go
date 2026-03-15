package paractice

import (
	"fmt"
	"sync"
)

func merge(ch1, ch2 <-chan int) <-chan int {
	res := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v1 := range ch1 {
			res <- v1
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v2 := range ch2 {
			res <- v2
		}
	}()

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}

func testMerge() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 1; i <= 5; i += 2 {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 2; i <= 6; i += 2 {
			ch2 <- i
		}
		close(ch2)
	}()

	merged := merge(ch1, ch2)

	for v := range merged {
		fmt.Println(v)
	}

	fmt.Println("programm finished")
}
