package stringmethods

import (
	"fmt"
	"strings"
)

func stringMethods() {
	/*
		Редактирование строк
	*/
	mySurnameWithMistake := "Sobolev"
	rightSurname := strings.Replace(mySurnameWithMistake, "l", "r", 1)
	fmt.Println(rightSurname)

	cats := "Котики важны. Котики улучшают эмоциональное здоровье."
    workout := strings.ReplaceAll(cats, "Котики", "Тренировки")

    fmt.Println(workout)
    fmt.Println(strings.Replace(cats, "Котики", "Тренировки", 1))

}
