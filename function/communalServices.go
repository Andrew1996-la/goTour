package function

import "fmt"

const (
	electricity = 5.88
	hotWater    = 156.28
	coldWater   = 32.39
	drainage    = 29.69
)

func payment(lastCounter, currentCounter int, tariff float64) float64 {
	return float64(currentCounter-lastCounter) * tariff
}

func electricityCalc(lastCounter, currentCounter int) float64 {
	return payment(lastCounter, currentCounter, electricity)
}

func hotWaterCalc(lastCounter, currentCounter int) float64 {
	return payment(lastCounter, currentCounter, hotWater)
}

func coldWaterCalc(lastCounter, currentCounter int) float64 {
	return payment(lastCounter, currentCounter, coldWater)
}

func waterDrainageCalc(coldWaterLast, coldWaterCurrent, hotWaterLast, hotWaterCurrent float64) float64 {
	drainageIndicator := (coldWaterCurrent - coldWaterLast) + (hotWaterCurrent - hotWaterLast)
	return drainageIndicator * drainage
}

func calculateTotal(payment ...float64) float64 {
	var total float64
	for _, value := range payment {
		total += value
	}
	return total
}

func communalServices() {
	paymentsElectocity := electricityCalc(13130, 13780)
	fmt.Println("В этом месяце платеж за электроэнергию составил", paymentsElectocity)

	paymentsHotWater := hotWaterCalc(57, 60)
	fmt.Println("В этом месяце платеж за горячую воду составил", paymentsHotWater)

	paymentsColdWater := coldWaterCalc(191, 199)
	fmt.Println("В этом месяце платеж за холодную воду составил", paymentsColdWater)

	paymentWaterDrainage := waterDrainageCalc(57, 60, 191, 199)
	fmt.Println("В этом месяце платеж за водоотведение составил", paymentWaterDrainage)

	payments := []float64{paymentsElectocity, paymentsHotWater, paymentsColdWater, paymentWaterDrainage}

	total := calculateTotal(payments...)

	fmt.Println("Всего придется заплатить", total)
}
