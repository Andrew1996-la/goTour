package processingErrors

import (
	"errors"
	"fmt"
	"strconv"
)

// Создаем переменную ошибки тут
var (
	ErrConvToFloat       = errors.New("ошибка преобразования строки в число с плавающей точкой")
	ErrNotPositiveNumber = errors.New("число должно быть больше нуля")
)

// Функция сложения
func addPositive(num1, num2 string) (float64, error) {
	firstNum, err := strconv.ParseFloat(num1, 64)
	if err != nil {
		return 0, ErrConvToFloat
	}
	secondNum, err := strconv.ParseFloat(num2, 64)
	if err != nil {
		return 0, ErrConvToFloat
	}

	if firstNum <= 0 || secondNum <= 0 {
		return 0, ErrNotPositiveNumber
	}

	return firstNum + secondNum, nil
}

func addPositiveNumber() {
	// ---------------------------- first case
	num1 := "5.2"
	num2 := "2.7"

	result, err := addPositive(num1, num2)

	if err != nil {
		err = fmt.Errorf("ошибка в ходе выполнения программы: %v", err)
		fmt.Println(err)
	} else {
		fmt.Println("результат сложения", result)
	}

	// ---------------------------- second case

	num1 = "two"
	num2 = "five"

	result, err = addPositive(num1, num2)
	if err != nil {
		err = fmt.Errorf("ошибка в ходе выполнения программы: %v", err)
		fmt.Println(err)
	} else {
		fmt.Println("результат сложения:", result)
	}

	// ---------------------------- third case

	num1 = "5.4"
	num2 = "0.0"

	result, err = addPositive(num1, num2)
	if err != nil {
		err = fmt.Errorf("ошибка в ходе выполнения программы: %v", err)
		fmt.Println(err)
	} else {
		fmt.Println("результат сложения:", result)
	}
}
