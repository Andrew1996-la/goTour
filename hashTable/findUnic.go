package hashtable

import "fmt"

func findUniqueNumber(nums []int) int {
	hashNums := make(map[int]int)

	// посчитаем сколько всего
	for _, num := range nums {
		hashNums[num]++
	}

	// поиск уникольного то есть
	for _, num := range nums {
		if hashNums[num] > 1 {
			continue
		}

		return num
	}

	return -1
}

func findUnicMain() {
	nums := []int{3, 1, 4, 1, 5, 3}

	unic := findUniqueNumber(nums)
	fmt.Println("Уникальное число:", unic)
}
