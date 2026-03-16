package paractice

import (
	"fmt"
	"time"
)

func timerWorker(jobs <-chan int, stop <-chan struct{}) {
	for {
		select {
		case num, ok := <-jobs:
			if !ok {
				fmt.Println("jobs was closed")
				return
			}
			fmt.Println(num)
		case <-stop:
			fmt.Println("time is over")
			return
		}
	}
}

func timerWorkerTest() {
	jobs := make(chan int)
	stop := make(chan struct{})

	go timerWorker(jobs, stop)

	go func() {
		for i := 1; i <= 100; i++ {
			time.Sleep(200 * time.Millisecond)
			jobs <- i
		}
		close(jobs)
	}()

	time.Sleep(2 * time.Second)

	close(stop) //то же что и stop <- struct{}{}

	time.Sleep(time.Second)

}
