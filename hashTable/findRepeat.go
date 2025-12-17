package hashtable

import "fmt"

func findRepeatNum(nums []int) int {
	numMap := make(map[int]bool)

	for _, num := range nums {
		if _, ok := numMap[num]; ok {
			return num
		} else {
			numMap[num] = true
		}
	}

	return -1
}

func any() {
	numList := []int{3, 1, 4, 1, 5, 3}

	repeat := findRepeatNum(numList)
	fmt.Println(repeat)
}
