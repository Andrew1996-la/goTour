package mutex_ex

import (
	"fmt"
	"sync"
	"time"
)

type Cache1 struct {
	mu sync.RWMutex
	m  map[int]int
}

func (cache *Cache1) Get(i int) int {
	cache.mu.RLock() // ставлю блок на взятие мьютекса
	v, ok := cache.m[i] // читаю не боясь изменений
	cache.mu.RUnlock() // разблокирую мьютекс
	if ok {
		return v
	}

	// получаем значение для указанного ключа
	cache.mu.Lock() // перед изменением ставлю мьютекс
	defer cache.mu.Unlock() // по окончании функции нужно снять мьютекс
	v = 2 * i
	cache.m[i] = v
	return v
}

func main() {
	cache := Cache1{
		m: make(map[int]int),
	}
	for i := 0; i < 20; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				cache.Get(j)
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(len(cache.m))
}
