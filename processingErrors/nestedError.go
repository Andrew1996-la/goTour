package processingErrors

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	ErrDivisionByZero = errors.New("division by zero")
	ErrConversion     = errors.New("conversion error")
)

func div(num1, num2 string) (int, error) {
	firstNum, err := strconv.Atoi(num1)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrConversion, err)
	}
	secondNum, err := strconv.Atoi(num2)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrConversion, err)
	}

	if secondNum == 0 {
		return 0, ErrDivisionByZero
	}

	return firstNum / secondNum, nil
}

func nestedError() {
	res, err := div("5", "0")
	if err != nil {
		if errors.Is(err, strconv.ErrSyntax) {
			fmt.Println("ошибка преобразования строки в целое")
		} else {
			fmt.Println("деление на 0")
		}
	} else {
		fmt.Println(res)
	}

	res, err = div("abc", "def")
	if err != nil {
		if errors.Is(err, strconv.ErrSyntax) {
			fmt.Println("ошибка преобразования строки в целое")
		} else {
			fmt.Println("деление на 0")
		}
	} else {
		fmt.Println(res)
	}
}
