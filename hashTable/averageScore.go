package hashtable

import "github.com/k0kubun/pp"

type Grade struct {
	Student string
	Subject string
	Score   int
}

/*
Посчитать средний балл по каждому предмету для каждого студента
Найти студента с самым высоким средним баллом
*/
func averageScore() {
	grades := []Grade{
		{"Anna", "Math", 90},
		{"Anna", "Math", 95},
		{"Anna", "CS", 100},
		{"Bob", "Math", 70},
		{"Bob", "CS", 80},
	}

	studentsMap := make(map[string][]int)
	averageScore := make(map[string]float64)
	var bestStudent string

	// собрать оценски в слайс для анализа
	for _, student := range grades {
		if _, ok := studentsMap[student.Student]; !ok {
			studentsMap[student.Student] = make([]int, 0)
		}
		studentsMap[student.Student] = append(studentsMap[student.Student], student.Score)
	}

	// поиск среднего
	var maxAvearage float64
	for name, listScore := range studentsMap {
		sum := 0
		countScore := len(listScore)

		for _, score := range listScore {
			sum += score
		}
		averageVal := float64(sum) / float64(countScore)

		if averageVal > maxAvearage {
			maxAvearage = averageVal
			bestStudent = name
		}

		averageScore[name] = averageVal
	}

	pp.Println(averageScore)
	pp.Println(bestStudent)
}
