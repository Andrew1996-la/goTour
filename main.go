package main

import "fmt"

func Average(nums ...int) int {
	sum := 0

	for _, v := range nums {
		sum += v
	}
	return sum / len(nums)
}

func main() {
	input := [][]int{
		{10, 6},
		{-5, 22, 89},
		{0, 78, 120, 22, 56},
		{6, 11, -9, 1, -9},
		{0, 0, 0, 0},
		{},
		{-7},
		{100345, 1345780, -67342, 782434, 2},
	}
	result := []int{8, 35, 55, 0, 0, 0, -7, 432243}

	for i, pars := range input {
		if len(pars) == 0{
			continue
		}

		v := Average(pars...)
		// закончите код теста
		if v != result[i] {
			fmt.Println(i, ":", v, "!=", result[i])
			return
		}

	}
	fmt.Println("OK")
}
