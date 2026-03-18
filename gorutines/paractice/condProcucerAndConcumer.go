package paractice

import (
	"fmt"
	"sync"
	"time"
)

// Горутина-произовдитель
func producerWithCond(
	data *[]int,
	cond *sync.Cond,
	mu *sync.Mutex,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)

		mu.Lock()
		*data = append(*data, i)
		fmt.Println("Добавили", i)
		cond.Signal()
		mu.Unlock()
	}
}

// Горутина-потребитель
func consumerWithCond(
	data *[]int,
	cond *sync.Cond,
	mu *sync.Mutex,
) {
	for {
		mu.Lock()
		for len(*data) == 0 {
			cond.Wait()
		}
		top := (*data)[0]
		*data = (*data)[1:]
		mu.Unlock()

		fmt.Println("Получили", top)
	}
}

func condProcucerAndConcumer() {
	var data []int
	mu := &sync.Mutex{}
	cond := sync.NewCond(mu)
	wg := &sync.WaitGroup{}

	go consumerWithCond(&data, cond, mu)

	wg.Add(1)
	go producerWithCond(&data, cond, mu, wg)

	wg.Wait()

	fmt.Println("Program finished")
}
