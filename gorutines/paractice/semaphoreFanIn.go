package paractice

import (
	"fmt"
	"sync"
)

/*
1. Обрабатывать задачи параллельно
2. Не более 3 одновременно (через semaphore)
3. Каждая задача считает num * num
4. Результаты складывать в канал results
5. В конце вывести все результаты
*/
func semaphoreFanIn() {
	sem := make(chan struct{}, 3)
	result := make(chan int)
	wg := &sync.WaitGroup{}

	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	go func() {
		for res := range result {
			fmt.Println(res)
		}
	}()

	for _, task := range tasks {
		sem <- struct{}{} // занял слот

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defer func() { <-sem }() // высвободил слот

			result <- i * i
		}(task)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	wg.Wait()

	fmt.Println("finish work")
}
