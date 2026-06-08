package multiple

import "testing"

func TestMultiple(t *testing.T) {
	a := 3
	b := 2

	res := Multiply(a, b)

	if res != 6 {
		t.Errorf("expected 5, got %d", res) // Тест упадет но тестирование продолжится
		// t.Fatalf("expected 6, got %d", res) // Тестирование прервется после падения
	}
}
