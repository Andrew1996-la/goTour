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

// startServer запускает HTTP сервер в отдельной горутине
func startServer(server *http.Server) {
	// ListenAndServe блокирует поток, пока сервер работает.
	// Если произошла ошибка (кроме ErrServerClosed, который возникает при Shutdown), выводим её
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("Ошибка сервера", err)
	}
}

// baseHttpGracefulShutdown демонстрирует корректное завершение HTTP сервера
func baseHttpGracefulShutdown() {
	// Роутинг: главная страница
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Главная страница") // ответ клиенту
		fmt.Println("Запрос на /")          // лог в консоль
	})

	// Роутинг: страница "Обо мне"
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Обо мне")       // ответ клиенту
		fmt.Println("Запрос на /about") // лог в консоль
	})

	// Настройка HTTP сервера
	server := &http.Server{
		Addr:    ":8080",              // порт сервера
		Handler: http.DefaultServeMux, // использует стандартный mux
	}

	// Создаём канал для получения сигналов ОС
	// Буфер 1 позволяет избежать блокировки, если сигнал придёт раньше чтения из канала
	quit := make(chan os.Signal, 1)

	// signal.Notify подписывает канал quit на сигналы SIGINT (Ctrl+C) и SIGTERM (стандартное завершение)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Запускаем сервер в отдельной горутине, чтобы main не блокировался
	go startServer(server)
	fmt.Println("Сервер успешно запущен")

	// Блокируем выполнение до получения сигнала завершения
	<-quit
	fmt.Println("Принят сигнал для завершения работы http сервера")

	// Создаём контекст с таймаутом 5 секунд для корректного завершения сервера
	// Это позволит завершить все активные запросы, не прерывая их сразу
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // обязательно вызываем cancel для освобождения ресурсов контекста

	// Вызываем Shutdown с контекстом
	// Shutdown безопасно завершает сервер, позволяя текущим запросам закончиться
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Ошибка при завершении сервера") // если что-то пошло не так
	} else {
		fmt.Println("Работа сервера завершена успешно") // сервер остановился корректно
	}
}
