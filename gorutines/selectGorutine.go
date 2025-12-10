package gorutines

import (
	"fmt"
	"strconv"
	"time"
)

func SelectGorutine() {

	intCh := make(chan int)
	strCh := make(chan string)

	go func() {
		i := 0

		for {
			intCh <- i
			i++

			time.Sleep(3 * time.Second)
		}
	}()

	go func() {
		i := 0

		for {
			strCh <- "Hello" + strconv.Itoa(i)
			i++

			time.Sleep(1 * time.Second)
		}
	}()

	/*
		Конструкция ниже будет исполняться долго

		Чтение из канала с числами происходит дольше
		чем чтение из канала со строками. Таким образом
		у нас просто подвисает основная горутина когда
		ждет значение число. Хотя горутина со строкой уже давно
		записана и только ждет когда основаная горутина освободится
		и придет прочитать.

		Не эффективно
	*/
	// for {
	// 	number := <- intCh
	// 	fmt.Println(number)
	// 	str := <- strCh
	// 	fmt.Println(str)
	// }

	/*
		конструкция select позволяет эффективно работать с горутинами

		то есть главная горутина смотрит в каком канале сейчас есть что читать
		и читает оттуда где все уже готово. Пришла строка а число еще нет? Значит
		прочитаем из канала с строкой. Если есть что читать в обоих каналах
		то прочитается любая случайная из доступных
	*/
	for {
		select {
		case number := <-intCh:
			fmt.Println(number)
		case str := <-strCh:
			fmt.Println(str)
		}
	}
}
