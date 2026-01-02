package task

// Проверить, есть ли элемент в слайсе.
func Contains[T comparable](arr []T, v T) bool {
	for _, a := range arr {
		if a == v {
			return true
		}
	}
	return false
}
