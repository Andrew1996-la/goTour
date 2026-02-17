package gracefulShutdown_

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func startServer(server *http.Server) {
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("Ошибка сервера", err)
	}
}

func baseHttpGracefulShutdown() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Главная страница")
		fmt.Println("Запрос на /")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Обо мне")
		fmt.Println("Запрос на /about")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: http.DefaultServeMux,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go startServer(server)
	fmt.Println("Сервер успешно запущен")

	<-quit
	fmt.Println("Принят сигнал для завершения работы http сервера")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Ошибка при завершении сервера")
	} else {
		fmt.Println("Работа сервера завершена успешно")
	}

}
