package gorutines

import (
	"fmt"
	"sync"
	"time"
)

func cond() {
	var data []int
	mu := &sync.Mutex{}
	cond := sync.NewCond(mu)

	// горутина потребитель
	go func() {
		for {
			mu.Lock()
			for len(data) == 0 {
				cond.Wait()
			}

			val := data[0]
			data = data[1:]

			mu.Unlock()
			fmt.Println("GOT:", val)
		}
	}()

	// горутина производитель
	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Second) // имитация работы

			mu.Lock()
			data = append(data, i)
			cond.Signal() // оповещаем потребителя что данные есть
			mu.Unlock()

			fmt.Println("Sent", i)
		}
	}()

	time.Sleep(7 * time.Second)
	fmt.Println("Program finished")
}
