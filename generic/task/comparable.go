package task

func IsEqual[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	}
	return false
}
