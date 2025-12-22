package gorutines

import (
	"fmt"
	"sync"
)

/*
атомики выполняются в одно действие в процессоре, поэтому можно избежать гонки с помощью них
но они работают только с числами
так же они значительно дольше работают

var counter atomic.Uint64
*/

var counter int
var mtx sync.Mutex // иницаилизируем mutex

/*
Мьютекс помогает избежать состояния гонки, когда несколько горутин
изменяют одно значение.
*/
func increase(wg *sync.WaitGroup) {
	defer wg.Done()
	for range 1000 {
		mtx.Lock()   // блокируем выполнение остальными горутнами исполненение пока взявшая mutex не разблокируется
		counter++    // значение безопасно менятес
		mtx.Unlock() // мьютекс разблокируется и его может взять другоая горутина
	}
}

func MutexRaceCondition() {
	wg := &sync.WaitGroup{}

	wg.Add(10)

	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)

	wg.Wait()

	fmt.Println(counter)
}
