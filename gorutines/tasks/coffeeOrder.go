package tasks

/*
Условия задачи

Клиенты
Клиенты приходят каждые 200 мс
Каждый клиент делает заказ:
	"espresso"
	"latte"
	"cappuccino"
	Клиенты отправляют заказы в канал ordersCh
	❗ Если очередь заказов переполнена, клиент уходит, заказ теряется

Баристы (2 горутины)
В кофейне работают 2 баристы
Бариста:
	читает заказ из ordersCh
	готовит кофе:
	espresso → 300 мс
	latte → 600 мс
	cappuccino → 800 мс
	после приготовления отправляет сообщение в readyCh
	Закрытие кофейни

Кофейня работает 3 секунды
После закрытия:
	новые клиенты не принимаются
	баристы доделывают текущие заказы
	программа завершается корректно
*/

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Client struct {
	Id    int
	Order string
}

func makeOrder() Client {
	menu := []string{
		"espresso",
		"cappuccino",
		"latte",
	}
	shooseRandomOrder := menu[rand.Intn(3)]

	id := rand.Int()

	return Client{
		Id:    id,
		Order: shooseRandomOrder,
	}
}

func barista(orders chan Client, ready chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for order := range orders {
		msg := fmt.Sprintln("Заказ", order.Order, "готов")

		switch order.Order {
		case "espresso":
			time.Sleep(300 * time.Millisecond)
		case "cappuccino":
			time.Sleep(800 * time.Millisecond)
		case "latte":
			time.Sleep(600 * time.Millisecond)
		}
		ready <- msg
	}
}

// Задача: Кофейня с баристой и заказами
func CoffeeOrder() {
	wg := &sync.WaitGroup{}

	ordersCh := make(chan Client, 10)
	readyCh := make(chan string, 10)

	go func() {
		time.Sleep(200 * time.Millisecond)
		timeForOrder := 0
		for {
			if timeForOrder >= 3000 {
				close(ordersCh)
				break
			}

			order := makeOrder()

			switch order.Order {
			case "espresso":
				timeForOrder += 300
			case "cappuccino":
				timeForOrder += 800
			case "latte":
				timeForOrder += 600
			}

			ordersCh <- makeOrder()
		}
	}()

	wg.Add(2)
	go barista(ordersCh, readyCh, wg)
	go barista(ordersCh, readyCh, wg)

	go func() {
		wg.Wait()
		close(readyCh)
	}()

	for mes := range readyCh {
		fmt.Println(mes)
	}

	fmt.Println("Кофеня закрылась")

}
