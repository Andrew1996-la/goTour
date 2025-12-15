package tasks

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type Trip struct {
	Driver string
	City   string
	Price  int
}

type DriverStats struct {
	Trips int
	Total int
	Avg   float64
}

func GetDriverInfo() {
	driversIncome := make(map[string]int)
	driversDirections := make(map[string]int)
	driversStatistic := make(map[string]DriverStats)

	var bestDriverName string
	var bestDriverSalary int

	trips := []Trip{
		{"Alex", "Moscow", 500},
		{"Bob", "Moscow", 700},
		{"Alex", "SPB", 400},
		{"Alex", "Moscow", 600},
		{"Bob", "SPB", 300},
		{"Bob", "SPB", 300},
		{"Chris", "Moscow", 1000},
	}

	for _, trip := range trips {
		driversIncome[trip.Driver] += trip.Price
		driversDirections[trip.City]++
	}

	for driverName, driverSalary := range driversIncome {
		if driverSalary > bestDriverSalary {
			bestDriverName = driverName
			bestDriverSalary = driverSalary
		}
	}

	for _, trip := range trips {
		stats := driversStatistic[trip.Driver]
		fmt.Println(stats)
		stats.Trips++
		stats.Total += trip.Price
		stats.Avg = float64(stats.Total) / float64(stats.Trips)
		driversStatistic[trip.Driver] = stats

	}

	pp.Println(driversIncome)
	pp.Println(driversStatistic)
	fmt.Println(bestDriverName, bestDriverSalary)
}
