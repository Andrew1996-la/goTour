package tasks

import "github.com/k0kubun/pp"


type UserAnalytics struct {
	Total int
	Avg   float64
}

func MapTask() {
	personOrderList := make(map[string][]Order)
	personStat := make(map[string]UserAnalytics)
	mostPopular := make(map[string]string)

	orders := []Order{
		{"Alice", "Book", 500},
		{"Alice", "Book", 200},
		{"Bob", "Pen", 100},
		{"Alice", "Laptop", 120000},
		{"Bob", "Notebook", 300},
		{"Alice", "Pen", 100},
		{"Charlie", "Book", 500},
		{"Bob", "Book", 500},
	}

	/*
		Собрать map пользователей → список их заказов

		Для каждого пользователя посчитать:
		общую сумму
		средний чек

		Найти самый популярный товар у каждого пользователя
	*/

	for _, order := range orders {
		if personOrderList[order.User] == nil {
			personOrderList[order.User] = make([]Order, 0)
		}
		personOrderList[order.User] = append(personOrderList[order.User], order)
	}

	for name, orderList := range personOrderList {
		var total int
		for _, order := range orderList {
			total += order.Price
		}
		avg := float64(total) / float64(len(orderList))
		personStat[name] = UserAnalytics{
			Total: total,
			Avg:   avg,
		}
	}

	for user, orderList := range personOrderList {
		var maxCount int
		var popularItem string

		products := make(map[string]int)

		for _, order := range orderList {
			products[order.Item]++
		}
		for item, count := range products {
			if count > maxCount {
				maxCount = count
				popularItem = item
			}
		}
		mostPopular[user] = popularItem

	}
	pp.Println(mostPopular)
}
