package gorutines

import (
	"fmt"
	"sync"
	"time"
)

func mechanic(item string, n int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		fmt.Printf("mechanic %d fix %s\n", n, item)
		time.Sleep(250 * time.Millisecond)
	}
}

func mechanicService() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go mechanic("Phone", 1, wg)
	wg.Add(1)
	go mechanic("Car", 2, wg)

	wg.Wait()

	fmt.Println("program is done")
}
