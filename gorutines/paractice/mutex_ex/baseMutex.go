package mutex_ex

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu sync.Mutex
	m  map[int]int
}

func (cache *Cache) Get(i int) int {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	v, ok := cache.m[i]
	if ok {
		return v
	}
	// получаем значение для указанного ключа
	v = 2 * i
	cache.m[i] = v
	return v
}

func baseMutexTest() {
	cache := Cache{
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
