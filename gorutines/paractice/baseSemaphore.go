package paractice

import (
	"fmt"
	"time"
)

func baseSemaphore() {
	sem := make(chan struct{}, 3)

	for i := 1; i <= 10; i++ {
		sem <- struct{}{} // Занимаю слот горутины

		go func(i int) {
			fmt.Println("start", i)
			time.Sleep(1 * time.Second)
			fmt.Println("finish", i)

			<-sem // Высвобождаю слот
		}(i)
	}

	time.Sleep(5 * time.Second)

	fmt.Println("program finished")
}
