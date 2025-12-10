package main

import (
	"fmt"
	"time"
)

// фукция по добыче угля
// принимает:
// - transferPoint (канал) - место куда нести уголь
// - номер похода
func mine(trancferPoint chan int, n int)  {
	fmt.Println("поход в шахту номер", n, "начался")
	time.Sleep(1 * time.Second) // событие по добыче угля
	fmt.Println("поход в шахту номер", n, "закончился")

	// каждый шахтер несет уголь в место приема угля главныму шахтеру
	trancferPoint <- 10
}

func main() {
	coal := 0

	// создается место приемки угля(канал) главного шахтера
	// от второстепенных шахтеров
	// канал не буфиризированный
	trancferPoint := make(chan int)

	// время начала работы
	start := time.Now()

	// делим рабоу на 3 паралельных потока
	go mine(trancferPoint, 1)
	go mine(trancferPoint, 2)
	go mine(trancferPoint, 3)

	// главный шахтер забирае уголь из места куда его
	// приносят второстепенные и складывает на склад
	coal += <- trancferPoint
	coal += <- trancferPoint
	coal += <- trancferPoint

	// время окончания работы
	timeForLoot := time.Since(start)

	fmt.Println("Всего шахтер добыл", coal, "угля за", timeForLoot)
}
