package task

import "fmt"

// Написать функцию, которая возвращает первый элемент слайса.
func First[T any](arr []T) T {
	fmt.Printf("Превый элемент слайста типа %T\n", arr[0])
	return arr[0]
}
