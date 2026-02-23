package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 1,
	}

	//	формирование запроса
	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080/time", nil)
	if err != nil {
		fmt.Println("ошибка формирования запроса", err)
		return
	}

	// исполнение запроса и получение результата
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("ошибка отправки запроса", err)
		return
	}

	// закрытие файла после чтения из body
	defer response.Body.Close()

	// запись результата в переменную
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("ошибка чтения тела запроса", err)
		return
	}

	fmt.Println(string(body))
}
