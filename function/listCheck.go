package function

import (
	"fmt"
	"math"
	"sort"
)

// Average возвращает среднее арифметическое элементов слайса []int.
// Напишите код функции
func Average(list []int) float64 {
	if len(list) == 0 {
		return 0
	}

	var sum float64
	for _, v := range list {
		sum += float64(v)
	}
	return sum / float64(len(list))
}

// Range возвращает размах числовой последовательности.
// Напишите код функции
func Range(list []int) int {
	if len(list) == 0 {
		return 0
	}

	if len(list) == 1 {
		return 0
	}

	minNum := list[0]
	maxNum := list[0]

	for _, v := range list {
		if v < minNum {
			minNum = v
		}
		if v > maxNum {
			maxNum = v
		}
	}
	return maxNum - minNum
}

func Median(list []int) int {
	if len(list) == 0 {
		return 0
	}

	sort.Ints(list)

	if len(list)%2 == 0 {
		midIndex := len(list) / 2
		return (list[midIndex-1] + list[midIndex]) / 2

	}

	return list[len(list)/2]
}

func Mode(list []int) (modeList []int, modeCount int) {
	if len(list) == 0 {
		return []int{}, 1
	}

	modeMap := make(map[int]int)
	res := make([]int, 0)
	maxRepeat := 1

	for _, v := range list {
		modeMap[v]++
	}

	for value, count := range modeMap {
		if count > 1 && count >= maxRepeat {
			res = append(res, value)
			maxRepeat = count
		}
	}

	sort.Ints(res)
	return res, maxRepeat
}

func listCheck() {
	lists := [][]int{
		{},
		{57},
		{78, -7},
		{99, 200, 0},
		{4, 4, 4, 3},
		{102, -7, 44, -7, 102},
		{82, -23, 1, 5, 98, 100},
		{100000, 90000, 20000, 20000, 20000, 22000, 25500, 22000},
	}

	averages := []float64{
		0, 57, 36, 100, 4, 47, 44, 39938,
	}
	ranges := []int{
		0, 0, 85, 200, 1, 109, 123, 80000,
	}
	medians := []int{
		0, 57, 35, 99, 4, 44, 43, 22000,
	}
	modes := [][]int{
		{}, {}, {}, {},
		{4},
		{-7, 102}, {},
		{20000},
	}
	mcount := []int{
		1, 1, 1, 1, 3, 2, 1, 3,
	}

	for i, list := range lists {
		if average := math.Round(Average(list)); average != averages[i] {
			fmt.Printf("average %d: %.2f != %.2f\n", i, averages[i], average)
		}
		if r := Range(list); r != ranges[i] {
			fmt.Printf("range %d: %d != %d\n", i, ranges[i], r)
		}
		if median := Median(list); median != medians[i] {
			fmt.Printf("median %d: %d != %d\n", i, medians[i], median)
		}
		mode, count := Mode(list)
		if len(mode) != len(modes[i]) {
			fmt.Printf("len mode %d: %v != %v'\n", i, modes[i], mode)
		} else {
			for j, v := range mode {
				if v != modes[i][j] {
					fmt.Printf("mode %d: %v != %v\n", i, modes[i], mode)
				}
			}
		}
		if count != mcount[i] {
			fmt.Printf("mcount %d: %d != %d\n", i, mcount[i], count)
		}
	}
	fmt.Println("Тестирование завершено")
}
