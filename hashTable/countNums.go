package hashtable

import "fmt"

func countNums(nums []int) map[int]int {
	numberStat := make(map[int]int)

	for _, num := range nums {
		numberStat[num]++
	}

	return numberStat
}

// Нужно посчитать, сколько раз встречается каждое число.
func countNumsFunc() {
	nums := []int{1, 2, 2, 3, 1, 2}
	res := countNums(nums)
	fmt.Println(res)
}
