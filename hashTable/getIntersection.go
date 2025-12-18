package hashtable

import "fmt"

func getIntersection(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)

	hashNums := make(map[int]bool)

	for _, val := range nums1 {
		hashNums[val] = true
	}

	for _, val := range nums2 {
		if _, ok := hashNums[val]; ok {
			result = append(result, val)
			delete(hashNums, val)
		}
	}

	return result
}

func getIntersectionMain() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}

	intersection := getIntersection(nums1, nums2)

	fmt.Println("Пересечение:", intersection)
}
