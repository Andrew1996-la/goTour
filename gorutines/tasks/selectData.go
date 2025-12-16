package tasks

import (
	"fmt"
	"strconv"
	"time"
)

func SelectData() {
	chString := make(chan string)
	chInt := make(chan int)

	// пишем строки
	go func() {
		index := 1
		for {
			time.Sleep(3 * time.Second)
			chString <- "string" + strconv.Itoa(index)
			index++
		}
	}()

	// пишем числа
	go func() {
		index := 1
		for {
			time.Sleep(1 * time.Second)
			chInt <- index
			index++
		}
	}()

	// читаем строки и числа по мере поступления
	for {
		select {
		case str := <-chString:
			fmt.Println(str)
		case integer := <-chInt:
			fmt.Println(integer)
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("нет данных")
		}
	}
}
