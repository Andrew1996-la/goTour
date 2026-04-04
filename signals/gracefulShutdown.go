package signals

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// worker обрабатывает задачи из канала
func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("worker %d started task: %s\n", id, job)

		time.Sleep(3 * time.Second) // имитация работы

		fmt.Printf("worker %d finished task: %s\n", id, job)
	}

	fmt.Printf("worker %d stopped\n", id)
}

func gracefulShutdownTest() {
	jobs := make(chan string, 10)
	wg := &sync.WaitGroup{}

	// контекст для сигналов (Ctrl+C, SIGTERM)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// --- HTTP handler ---
	mux := http.NewServeMux()
	mux.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "cannot read body", http.StatusBadRequest)
			return
		}

		task := string(body)

		select {
		case jobs <- task:
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("task accepted"))
		default:
			http.Error(w, "queue is full", http.StatusServiceUnavailable)
		}
	})

	// --- запускаем воркеры ---
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, wg)
	}

	// --- создаём сервер ---
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// --- запускаем сервер в горутине ---
	go func() {
		fmt.Println("server started on :8080")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("server error:", err)
		}
	}()

	// --- ждём сигнал ---
	<-ctx.Done()
	fmt.Println("\nshutting down...")

	// --- даём 5 секунд на graceful shutdown ---
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1. останавливаем HTTP сервер (перестаёт принимать новые запросы)
	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Println("shutdown error:", err)
	}

	// 2. закрываем канал (воркеры дорабатывают оставшиеся задачи)
	close(jobs)

	// 3. ждём завершения воркеров
	wg.Wait()

	fmt.Println("graceful shutdown complete")
}
