package paractice

import (
	"fmt"
	"sync"
)

// fan-in это когда много источников(channels), но один примемник
func fanInWorker(res chan<- int, in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for n := range in {
		res <- n * 2
	}
}

func fanInTest() {
	resultCh := make(chan int)
	wg := &sync.WaitGroup{}

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		if i == 1 {
			go fanInWorker(resultCh, ch1, wg)
		}

		if i == 2 {
			go fanInWorker(resultCh, ch2, wg)
		}

		if i == 3 {
			go fanInWorker(resultCh, ch3, wg)
		}
	}

	go func() {
		for i := 1; i <= 5; i++ {
			ch1 <- i
			ch2 <- i
			ch3 <- i
		}

		close(ch1)
		close(ch2)
		close(ch3)
	}()

	// Необходимо ждать завершения горутин в отдельной горутине поскольку у нас несколько
	// писателей в канал и мы не знаем точно когда писатели закончат
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for n := range resultCh {
		fmt.Println(n)
	}
}
