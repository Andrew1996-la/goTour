package syncmap

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	store sync.Map
}

func NewCache() *Cache {
	return &Cache{
	}
}

func (c *Cache) Set(key string, value int, d time.Duration) {
	c.store.Store(key, value)
	timer := time.AfterFunc(d, func() {
		c.store.Delete(key)
	})
	timer.Reset(d)
}

func (c *Cache) Get(key string) (any, bool) {
	return c.store.Load(key)
}

func syncMapTest() {
	cache := NewCache()

	cache.Set("1", 567, 1000*time.Millisecond)
	cache.Set("2", 22, 2000*time.Millisecond)
	cache.Set("3", 9, 3000*time.Millisecond)

	print := func() {
		for _, key := range []string{"1", "2", "3"} {
			v, ok := cache.Get(key)
			if ok {
				fmt.Print(v, " ")
			}
		}
		fmt.Println()
	}

	print()
	time.Sleep(1500 * time.Millisecond)
	print()
	time.Sleep(1000 * time.Millisecond)
	print()
}
