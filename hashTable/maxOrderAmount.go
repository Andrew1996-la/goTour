package hashtable

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type Order struct {
	User   string
	Amount int
}

/*
Посчитать общую сумму покупок по каждому пользователю
Найти пользователя с максимальной суммой покупок
*/
func maxOrderAmount() {
	orders := []Order{
		{"Alice", 100},
		{"Bob", 200},
		{"Alice", 50},
		{"Bob", 70},
		{"Charlie", 300},
	}
	var maxAmount int
	var maxAmountUserName string

	orderStat := make(map[string]int, len(orders))

	for _, order := range orders {
		orderStat[order.User] += order.Amount
		if order.Amount > maxAmount {
			maxAmount = order.Amount
			maxAmountUserName = order.User
		}
	}

	pp.Println(orderStat)
	fmt.Println("Пользователь с наибольшей суммой заказов:", maxAmountUserName)
}
