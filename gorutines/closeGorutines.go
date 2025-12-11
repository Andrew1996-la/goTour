package gorutines

import (
	"fmt"
	"math/rand"
	"time"
)

func gorutines() {
	coal := 0

	transferPoint := make(chan int)

	// функция похода в шахту
	go func() {
		iterator := 3 + rand.Intn(4)
		fmt.Println("в шахте угля на", iterator, "походов")

		for i := 1; i <= iterator; i++ {
			transferPoint <- 10

			time.Sleep(1 * time.Second)
		}

		// после выработки угля нужно доложить начальнику что
		// угля больше нет. Закрываем канал
		close(transferPoint)
	}()

	/*
		for {
			// получаем значение и статус получения
			// если канал закрыт то получим в val дефолтное значение типа данных
			// а в статусе ok
			val, ok := <-transferPoint

			// проверяем закрыт ли канал
			if !ok {
				fmt.Println("Шахты была выроботана и канал закрыт")
				break
			}

			coal += val
		}
	*/

	// более распространенный способ работы с каналами в цикле
	// цикл прекращает работу сам когда канал закрывается
	for v := range transferPoint {
		coal += v
	}

	fmt.Println("Угля добыто", coal)
}
