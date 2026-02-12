package gorutines

import (
	"fmt"
	"strconv"
	"time"
)

func testSelect() {
	intChan := make(chan int)
	strChan := make(chan string)

	go func() {
		i := 0
		for {
			intChan <- i
			i++
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		i := 0
		for {
			strChan <- "hello " + strconv.Itoa(i)
			i++
			time.Sleep(1 * time.Second)
		}
	}()

	// позволяет обрабатывать значение из той горутины которая готова на данный момент.
	for {
		select {
		case intRes := <-intChan:
			fmt.Println(intRes)
		case strRes := <-strChan:
			fmt.Println(strRes)
		}
	}
}
