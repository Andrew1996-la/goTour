package anytask

import "fmt"

/*
дано:
массив №1 [1,3,5,7]
массив №2 [2,4,6,8]
Результат
массив [1,2,3,4,5,6,7,8]
*/
func TwoPointers() {
	arr1 := []int{1, 3, 5, 7}
	arr2 := []int{2, 4, 6, 8}

	res := make([]int, 0, len(arr1)+len(arr2))

	pointer1 := 0
	pointer2 := 0

	for pointer1 < len(arr1) && pointer2 < len(arr2) {
		if arr1[pointer1] < arr2[pointer2] {
			res = append(res, arr1[pointer1])
			pointer1++
		} else {
			res = append(res, arr2[pointer2])
			pointer2++
		}
	}

	for pointer1 < len(arr1) {
		res = append(res, arr1[pointer1])
		pointer1++
	}
	for pointer2 < len(arr2) {
		res = append(res, arr2[pointer2])
		pointer2++
	}

	fmt.Println(res)
}
