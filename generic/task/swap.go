package task

func Swap[T any](value1, value2 T) (T, T) {
	return value2, value1
}
