package slices

import "fmt"

func deleteFromSliceByIndex(sliceInt []int, deleteIndex int) []int {
	tmp := make([]int, len(sliceInt))
	copy(tmp, sliceInt)

	return append(tmp[:deleteIndex], tmp[deleteIndex+1:]...)
}

func deleteFromSliceByIndexTest() {
	delIndex := 2
	nums := []int{10, 20, 30, 40, 50, 60}

	res := deleteFromSliceByIndex(nums, delIndex)

	fmt.Println(res)
}
