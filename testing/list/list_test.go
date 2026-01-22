package list

import (
	"testing"
)

func TestMin(t *testing.T) {
	// варианты последовательностей
	// первая цифра массива это контрольное значение, остальные проверочные
    lists := map[string][]int{
        "three": {-50, 5, -50, 67},
        "four":  {-7, -7, 1, 43, 100},
        "two":   {3, 3, 4},
        "one":   {10, 10},
        "empty": {0},
    }

	// ожидаемые значения для каждой последовательности
   for name, list := range lists {
        t.Run(name, func(t *testing.T) {
            if Min(list[1:]) != list[0] {
                t.Errorf("get %d != want %d", Min(list[1:]), list[0])
            }
        })
    }
}

 