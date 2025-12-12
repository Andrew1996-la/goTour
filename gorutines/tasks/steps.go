package tasks

import (
	"fmt"
	"sync"
	"time"
)

func runs(person string, speed int, runningInfoCh chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	finishStep := 10

	for i := 1; i <= finishStep; i++ {
		str := fmt.Sprintf("i am %s. i do my %d step\n", person, i)
		runningInfoCh <- str
		time.Sleep(time.Duration(speed) * time.Millisecond)
	}
}

// Параллельный обработчи
func Steps() {
	wg := &sync.WaitGroup{}
	runningInfoCh := make(chan string)

	wg.Add(3)
	go runs("Warrior", 300, runningInfoCh, wg)
	go runs("Mage", 500, runningInfoCh, wg)
	go runs("Rouge", 200, runningInfoCh, wg)

	go func() {
		wg.Wait()
		close(runningInfoCh)
	}()

	for res := range runningInfoCh {
		fmt.Println(res)
	}

	fmt.Println("Programm finished")
}
