package tasks

import "github.com/k0kubun/pp"

type Order struct {
	User  string
	Item  string
	Price int
}

type Buyer struct {
	Name string
	Sum  int
}

func OrderList() {
	ordersMap := make(map[string]int)
	ordersCount := make(map[string]int)
	bigMap := make(map[string]map[string]int)

	bestBuyer := Buyer{}

	orders := []Order{
		{"Alice", "Book", 600},
		{"Bob", "Pen", 100},
		{"Alice", "Pen", 100},
		{"Alice", "Notebook", 300},
		{"Bob", "Book", 500},
		{"Charlie", "Book", 500},
		{"Bob", "Pen", 100},
	}

	// Посчитать общую сумму покупок по каждому пользователю
	// Посчитать, сколько раз каждый товар покупали
	// Найти пользователя с максимальной суммой покупок

	for _, order := range orders {
		ordersMap[order.User] += order.Price
		ordersCount[order.Item] += 1
	}

	for name, sum := range ordersMap {
		if sum > bestBuyer.Sum {
			bestBuyer.Name = name
			bestBuyer.Sum = sum
		}
	}

	for _, order := range orders {
		if bigMap[order.User] == nil {
			bigMap[order.User] = make(map[string]int)
		}

		bigMap[order.User][order.Item] += order.Price
	}

	pp.Println(ordersMap)
	pp.Println(ordersCount)
	pp.Println(bestBuyer)
	pp.Println(bigMap)
}
