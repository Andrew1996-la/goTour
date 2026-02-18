package gracefulShutdown_

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	endpoint              = "/ping"
	clientAnswer          = "pong"
	host                  = "localhost:8080"
	gracefulShutdownTimer = time.Second * 10
)

func handlePing(w http.ResponseWriter, r *http.Request) {
	log.Println("Получен запрос:", r.URL.Path)

	time.Sleep(time.Millisecond * 3000)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(clientAnswer))

	log.Println("Ответ обработан")
}

func startLocalServer(srv *http.Server) {
	err := srv.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Println("Ошибка сервера")
	}
}

func gracefulShutdown(srv *http.Server, sig os.Signal) error {
	log.Println("От ОС получен сигнал завершения программы", sig)

	ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdownTimer)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("Работа окончена таймаутом graceful-shutdown")
	}
	return nil
}

func gShoutDownTest() {
	mux := http.NewServeMux()

	mux.HandleFunc(endpoint, handlePing)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	server := &http.Server{
		Addr:    host,
		Handler: mux,
	}

	go startLocalServer(server)

	// главная горутина сдесь остановится и будет ожидать сигнал
	// код в main продолжит выполнение после получения сигнала
	sig := <-sigs

	err := gracefulShutdown(server, sig)
	if err != nil {
		log.Fatal("Сервер завершен с ошибкай", err)
	}

	log.Println("Сервер завершен безопасно")
}
