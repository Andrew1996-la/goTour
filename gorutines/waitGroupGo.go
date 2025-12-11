package gorutines

import (
	"fmt"
	"sync"
	"time"
)

func postman(wg *sync.WaitGroup, magazine string) {
	for i := 1; i <= 3; i++ {
		fmt.Println("Я почтальон и отнес газету", magazine, "в дом номер", i)
		time.Sleep(1 * time.Second)
	}

	// сигнализируем главному потоку что завершена горутина
	wg.Done()
}

func WaitGroupGo() {
	// создание waitGroup. Работаем по указателю
	wg := &sync.WaitGroup{}

	// фиксируем что начинается горутина
	wg.Add(1)
	go postman(wg, "Новости")

	// фиксируем что начинается горутина
	wg.Add(1)
	go postman(wg, "Игромания")

	/*
		поскольку мы уже просигнализировали что горутины были
		запущены. Мы должны дождаться их выполнения
	*/
	wg.Wait()
	fmt.Println("main завершился")
}
