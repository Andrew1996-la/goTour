package main

import "fmt"

// напишите рекурсивную функцию GCD, которая вычисляет НОД
// двух чисел по алгоритму Евклида
func GCD(n, m int) int {
	if n <= 0 || m <= 0 {
		return 0
	}

	r := n % m

	if r == 0 {
		return m
	}

	return GCD(m, r)
}

func test() {
	pairs := [][]int{
		{0, 5, 0},
		{7, -350, 0},
		{1460, 124, 4},
		{198, 42, 6},
		{1386, 4494, 42},
		{1470, 1575, 105},
	}
	for _, pair := range pairs {
		gcd := GCD(pair[0], pair[1])
		if gcd != pair[2] {
			fmt.Printf("НОД(%d, %d): %d != %d\n", pair[0], pair[1], gcd, pair[2])
		}
	}
	fmt.Println("Тестирование завершено")
}
