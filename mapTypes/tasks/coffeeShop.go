package tasks

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type Drink struct {
	Barista string
	Type    string
	Price   int
}

type BaristaStats struct {
	DrinksCount int
	TotalIncome int
	AvgPrice    float64
}

/*
Посчитать статистику по каждому бариста
Найти самого прибыльного бариста
Найти самый популярный напиток
*/
func CoffeeShop() {
	baristaMap := make(map[string]BaristaStats)
	popularDinks := make(map[string]int)

	var bestBarista string
	var mostPopularDrink string
	var countDrink int

	drinks := []Drink{
		{"Alice", "espresso", 150},
		{"Bob", "latte", 250},
		{"Alice", "latte", 250},
		{"Alice", "espresso", 150},
		{"Bob", "cappuccino", 300},
		{"Bob", "latte", 250},
		{"Alice", "cappuccino", 300},
	}

	for _, drinks := range drinks {
		barista := baristaMap[drinks.Barista]
		barista.DrinksCount++
		barista.TotalIncome += drinks.Price
		barista.AvgPrice = float64(barista.TotalIncome) / float64(barista.DrinksCount)

		baristaMap[drinks.Barista] = barista

		popularDinks[drinks.Type]++
	}

	var baristaResult int
	for name, baristaInfo := range baristaMap {
		if baristaInfo.TotalIncome > baristaResult {
			bestBarista = name
			baristaResult = baristaInfo.TotalIncome
		}
	}
	pp.Println(baristaMap)
	fmt.Println("Самый прибыльный бариста:", bestBarista)

	for drink, count := range popularDinks {
		if countDrink < count {
			mostPopularDrink = drink
			countDrink = count
		} 
	}

	fmt.Println("Самый популярный напиток:", mostPopularDrink)
}
