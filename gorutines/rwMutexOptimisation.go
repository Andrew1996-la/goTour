package gorutines

import (
	"fmt"
	"sync"
	"time"
)

var likes int
var mtx sync.RWMutex

/*
RWMutex позволяет паралельно читать значения если в момент чтения
горутины не изменяют текущюю переменную

таким образом если мы много читаем то код будет выполняться быстрее
*/

func setLike(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 100000; i++ {
		// базовая блокировка на запись
		mtx.Lock()
		likes++
		mtx.Unlock()
	}
}

func getLike(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 100000; i++ {
		// умная блокировка на чтение
		mtx.RLock()
		_ = likes
		mtx.RUnlock()
	}
}

func RWMutexOptimistion() {
	wg := &sync.WaitGroup{}

	timeStart := time.Now()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go setLike(wg)
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go getLike(wg)
	}

	wg.Wait()
	fmt.Println("Время выполнения:", time.Since(timeStart))
}
