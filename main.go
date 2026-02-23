package main

import (
	"fmt"
	"net/http"
	"time"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	var out string
	if r.URL.Path == "/time" || r.URL.Path == "/time/" {
		out = time.Now().Format("2006-01-02 15:04:05")
	} else {
		out = fmt.Sprintf("\nMethod: %s\nPath: %s\nHost: %s\n", r.Method, r.URL.Path, r.Host)
	}
	w.Write([]byte(out))
}

func main() {
	fmt.Println("Запускаем приложение")
	http.HandleFunc("/", mainHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Завершение работы")
}
