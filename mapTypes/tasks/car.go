package tasks

import "fmt"

// Посчитать, сколько раз каждая машина заезжала
// Найти самую часто встречающуюся машину
// Разделить машины на группы: ключ — марка машины значение — список индексов, на которых она встречается
func Car() {
	autoService := make(map[string]int)
	groupCar := make(map[string][]int)

	var mostPopularCar string
	var howOfthen int

	cars := []string{
		"BMW",
		"Audi",
		"BMW",
		"Tesla",
		"Audi",
		"BMW",
		"Tesla",
	}

	for index, car := range cars {
		autoService[car]++

		groupCar[car] = append(groupCar[car], index)
	}

	for auto, count := range autoService {
		if count > howOfthen {
			mostPopularCar = auto
			howOfthen = count
		}
	}

	fmt.Printf("Самая популярная марка в нашем автосервисе %s\nВстречается %d в неделю\n", mostPopularCar, howOfthen)
	fmt.Println(groupCar)
}
