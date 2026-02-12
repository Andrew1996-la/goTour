package gorutines

import (
	"fmt"
	"math/rand"
)

func closeChanel() {
	transferPoint := make(chan int)
	total := 0

	go func() {
		iteration := 3 + rand.Intn(4)
		fmt.Println("угля в шахте на", iteration, "походов")

		for i := 0; i < iteration; i++ {
			transferPoint <- 10
		}

		// закрытие канала
		close(transferPoint)
	}()

	// когда идет чтение, то цикл сам определяет что канал закрыт и не возникает deadlock
	for coal := range transferPoint {
		total += coal
	}

	fmt.Println(total)

}
