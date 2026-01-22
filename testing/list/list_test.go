package list

import "testing"

func TestMin(t *testing.T) {
	// варианты последовательностей
	lists := [][]int{
		{5, -50, 67},
		{-7, 1, 43, 100},
		{3, 4},
		{10},
		{},
	}

	// ожидаемые значения для каждой последовательности
	mins := []int{-50, -7, 3, 10, 0}

	for i, list := range lists {
		if Min(list) != mins[i] {
			t.Error(i, ":", Min(list), "!=", mins[i])
		}
	}
}
