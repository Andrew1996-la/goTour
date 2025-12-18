package hashtable

import "fmt"

func checkDublicate(nums []int) bool {
	hashNums := make(map[int]bool)

	for _, num := range nums {
		if _, ok := hashNums[num]; ok {
			return true
		}

		hashNums[num] = true
	}

	return false
}

func checkDoble() {
	nums := []int{1, 2, 3, 5}

	hasDouble := checkDublicate(nums)

	fmt.Println(hasDouble)
}
