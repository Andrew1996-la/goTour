package gorutines

import (
	"fmt"
	"time"
)

// фукция по добыче угля
// принимает:
// - transferPoint (канал) - место куда нести уголь
// - номер похода
func mine(trancferPoint chan int, n int) {
	fmt.Println("поход в шахту номер", n, "начался")
	time.Sleep(1 * time.Second) // событие по добыче угля
	fmt.Println("поход в шахту номер", n, "закончился")

	// каждый шахтер несет уголь в место приема угля главныму шахтеру
	// в не буфиризированном канале до тех пор пока другая горутина не
	// заберет значение из этой, код ниже trancferPoint <- 10 не будет исполняться
	trancferPoint <- 10
	fmt.Println("шахтер с похода", n, "передал уголь")
}

func GorutinesExample() {
	coal := 0

	// создается место приемки угля(канал) главного шахтера
	// от второстепенных шахтеров
	// канал буфиризированный резервируем для буфера 3 ячейки
	// шахтеры принесут уголь и оставят в буфере, а главный шахтер
	// придет за ним по возможности
	trancferPoint := make(chan int, 3)

	// время начала работы
	start := time.Now()

	// делим рабоу на 3 паралельных потока
	go mine(trancferPoint, 1)
	go mine(trancferPoint, 2)
	go mine(trancferPoint, 3)

	// главный шахтер забирае уголь из места куда его
	// приносят второстепенные и складывает на склад
	// перед каждым походом в канал спит немного)
	time.Sleep(3 * time.Second)
	coal += <-trancferPoint

	time.Sleep(3 * time.Second)
	coal += <-trancferPoint

	time.Sleep(3 * time.Second)
	coal += <-trancferPoint

	// время окончания работы
	timeForLoot := time.Since(start)

	fmt.Println("Всего шахтер добыл", coal, "угля за", timeForLoot)
}
