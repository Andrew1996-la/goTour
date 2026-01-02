package generic

import "fmt"

/*
Дженерики позволяют взаимодейстововать с функцией
разными типами данных. Тут например int и float64

дженерики обязательно нужно объявлять с возможными типами
иначе с ними не возможно будет взаимподействовать

для описания типов можно использовать кастомный constaint
он нужен для переиспользования и более чистого кода
*/

type Number interface {
	int | int64 | float64
}

func sum[T Number](a, b T) T {
	return a + b
}

/*
Позволяет так же создвать структуры с различными типами
*/
type Box[T any] struct {
	value T
}

func generics() {
	i := Box[int]{value: 1}
	s := Box[string]{value: "hello"}
	fmt.Println(i, s)
}
