package tasks

import (
	"fmt"
	"sync"
)

/*
Запустить горутины, каждая из которых:

	принимает одно число
	считает его квадрат

# Собрать результаты в один срез results
Гарантировать, что main:

	дождётся завершения всех горутин
	корректно выведет результат

Порядок результатов не важен.
*/

func doubel(n int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- n * 2
}

func doubelMain() {
	nums := []int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	for _, num := range nums {
		wg.Add(1)
		go doubel(num, ch, wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for doublVal := range ch {
		fmt.Println(doublVal)
	}
}
