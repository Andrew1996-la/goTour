package tasks

import (
	"fmt"
	"sync"
	"time"
)

// Задача: «Потоковое чтение файла через горутины»
func producer(ch chan string, str []string) {

	for _, v := range str {
		fmt.Println("Кладу в буфер строку: ", v)
		ch <- v
	}
	fmt.Println("Все строки переданы. Закрываем канал")

	close(ch)
}

func concumer(ch chan string, wg *sync.WaitGroup) {
	for line := range ch {
		fmt.Println("получено:", line)
		time.Sleep(1 * time.Second)
	}

	wg.Done()
}

func ProducerConcumer() {
	ch := make(chan string, 3)
	wg := &sync.WaitGroup{}

	lines := []string{
		"строка 1",
		"строка 2",
		"строка 3",
		"строка 4",
		"строка 5",
		"строка 6",
	}

	go producer(ch, lines)
	wg.Add(1)
	go concumer(ch, wg)
	wg.Add(1)
	go concumer(ch, wg)

	wg.Wait()
	fmt.Println("Завершение программы")
}
