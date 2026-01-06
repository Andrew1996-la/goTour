package processingErrors

import (
	"errors"
	"fmt"
)

var ErrCritical = errors.New("critical error")

func divider(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("%w, can't divide %d by zero ", ErrCritical, b)
	}

	return a / b, nil
}

func unwrapError() {
	result, err := divider(5, 0)
	if err != nil {
		fmt.Println(err)                // получить полную ошибку
		fmt.Println(errors.Unwrap(err)) // получить первоначальную ошбку
	} else {
		fmt.Println(result)
	}
}
