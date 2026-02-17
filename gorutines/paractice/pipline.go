package paractice

import (
	"fmt"
)

func generator(source chan<- int) {
	for i := 1; i <= 10; i++ {
		source <- i
	}
	close(source)
}

func processor(source <-chan int, result chan<- int) {
	for v := range source {
		result <- v * v
	}

	close(result)
}

func aggregator(result <-chan int, sumCh chan<- int) {
	sum := 0
	for v := range result {
		sum += v
	}
	sumCh <- sum

	close(sumCh)
}

func piplineTest() {
	source := make(chan int)
	result := make(chan int)
	sumCh := make(chan int)

	go generator(source)
	go processor(source, result)
	go aggregator(result, sumCh)

	sum := <-sumCh
	fmt.Println("sum:", sum)
	fmt.Println("finish pipline")
}
