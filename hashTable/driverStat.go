package hashtable

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type Trip struct {
	Driver string
	CarID  string
	Price  int
}

type PersonStat struct {
	total  int
	topCar string
}

/*
Посчитать общий доход каждого водителя
Для каждого водителя определить самую прибыльную машину
Найти водителя с максимальным доходом
*/
func driverStat() {
	trips := []Trip{
		{"Ivan", "Жигули", 30},
		{"Ivan", "Жигули", 150},
		{"Petr", "Мерседес", 200},
		{"Ivan", "Шкода", 50},
		{"Petr", "Мерседес", 100},
		{"Anna", "Жигули", 100},
	}
	stats := make(map[string]PersonStat)

	calcIncome := make(map[string]map[string]int)

	for _, trip := range trips {
		//понять если ли такой водитель
		if _, ok := calcIncome[trip.Driver]; !ok {
			//если нет создаем ему мапу
			calcIncome[trip.Driver] = make(map[string]int)

		}
		calcIncome[trip.Driver][trip.CarID] += trip.Price
	}

	// определить самую доходную машину водителя
	for driver, cars := range calcIncome {
		var total int
		var topCar string
		var maxCarIncome int

		for car, income := range cars {
			total += income

			if income > maxCarIncome {
				maxCarIncome = income
				topCar = car
			}
		}

		stats[driver] = PersonStat{
			total:  total,
			topCar: topCar,
		}
	}
	var maxTotal int
	var bestDriver string

	for name, carStat := range stats {
		if carStat.total > maxTotal {
			maxTotal = carStat.total
			bestDriver = name
		}
	}

	pp.Println(stats)
	fmt.Println("Лучший водитель", bestDriver)
}
