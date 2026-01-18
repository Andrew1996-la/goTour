package task

type NumberType interface {
	int | float64 | int64
}

func Sum[T NumberType](value1, value2 T) T {
	return value1 + value2
}
