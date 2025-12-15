package condition

import "fmt"

/*
Ставка 22% — на часть дохода свыше 50 млн ₽
Ставка 20% — на часть дохода свыше 20 и до 50 млн ₽ включительно
Ставка 18% — на часть дохода свыше 5 и до 20 млн ₽ включительно
Ставка 15% — на часть дохода свыше 2,4 и до 5 млн ₽ включительно
Ставка 13% — на часть дохода до 2,4 млн ₽ включительно
*/

const (
	Limit20 = 50_000_000
	Limit18 = 20_000_000
	Limit15 = 5_000_000
	Limit13 = 2_400_000
)

func ProgressivTask() {
	income := 57_000_000.0 // доход
	tax := 0.0             // налог

	recalculate := income

	// свыше 50 млн
	if recalculate > Limit20 {
		tax += (income - Limit20) * 22 / 100
		recalculate = Limit20
	}
	// от 20 до 50
	if recalculate > Limit18 && recalculate <= Limit20 {
		tax += (recalculate - Limit18) * 20 / 100
		recalculate = Limit18
	}
	// от 5 до 20
	if recalculate > Limit15 && recalculate <= Limit18 {
		tax += (recalculate - Limit15) * 18 / 100
		recalculate = Limit15
	}

	// от 2.4 до 5
	if recalculate > Limit13 && recalculate <= Limit15 {
		tax += (recalculate - Limit13) * 15 / 100
		recalculate = Limit13
	}
	// до 2.4
	if recalculate <= Limit13 {
		tax += recalculate * 13 / 100
	}

	fmt.Printf("Доход: %.2f, НДФЛ: %.2f\n", income, tax)
}
