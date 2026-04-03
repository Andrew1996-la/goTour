package signals

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func sigintSignal() {
	fmt.Println("ProcessID:", os.Getpid())

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)

	go func() {
		for {
			s := <-sigCh
			fmt.Println("got signal:", s)
		}
	}()

	time.Sleep(1 * time.Hour)
}
