package hashtable

import "fmt"

func twoSum(nums []int, target int) []int {
	hashNum := make(map[int]int)

	for index, num := range nums {
		need := target - num

		if prevIndex, ok := hashNum[need]; ok {
			return []int{prevIndex, index}
		}

		hashNum[num] = index

	}

	fmt.Println("Решение не найдено")
	return nil
}

func twoSumMain() {
	nums := []int{2, 7, 11, 15}
	target := 9

	sum := twoSum(nums, target)

	fmt.Println(sum)
}
