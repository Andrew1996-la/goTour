package main

import (
	"fmt"
	"goTour/models"
)

func main() {
	/*
		 Подключаю из пакета с моделями структуру Laptop
		для создания переменной с определенными полями
	*/
	mac := models.Laptop{
		Brand:  "Apple",
		Model:  "MacBook Air",
		Price:  1999.99,
		Memory: 16,
	}

	fmt.Printf("Laptop: %s %s with %d RAM, Price: $%.2f\n", mac.Brand, mac.Model, mac.Memory, mac.Price)
}
