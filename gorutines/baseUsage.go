package gorutines

import (
	"fmt"
	"time"
)

func mine(n int, coalCh chan int) {
	fmt.Printf("mine: %d start working\n", n)
	time.Sleep(time.Second * 1)
	fmt.Printf("mine: %d finished\n", n)

	coalCh <- 10
}

func baseUsage() {
	coalCh := make(chan int)
	totalCoal := 0

	go mine(1, coalCh)
	go mine(2, coalCh)
	go mine(3, coalCh)

	totalCoal += <-coalCh
	totalCoal += <-coalCh
	totalCoal += <-coalCh

	fmt.Printf("total count: %d\n", totalCoal)
}
