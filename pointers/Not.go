package pointers

import "fmt"

/*
Дан двумерный массив элементов из 0 и 1. Напишите функцию Not(),
которая проходит по всем элементам массива и меняет 0 на 1, а 1 на 0.
*/
func Not(arr *[5][10]int) {
	for idxStr, str := range *arr {
		for idxVal := range str {
			if arr[idxStr][idxVal] == 1 {
				arr[idxStr][idxVal] = 0
			} else {
				arr[idxStr][idxVal] = 1
			}
		}
	}
}

func NotTest() {
	arr := [5][10]int{
		{1, 1, 1, 0, 1, 1, 0, 1, 1, 1},
		{1, 1, 1, 0, 1, 1, 0, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 0, 1, 1, 0, 1, 1, 1},
		{1, 1, 1, 0, 1, 1, 0, 1, 1, 1},
	}
	Not(&arr)
	for _, v := range arr {
		fmt.Println(v)
	}
}
