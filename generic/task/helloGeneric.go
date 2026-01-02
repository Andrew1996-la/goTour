package task

import "fmt"

//Написать функцию, которая возвращает то, что ей передали.

func Identifity[T any](t T) T {
	fmt.Printf("тип аргумента %T\n", t)
	return t
}
