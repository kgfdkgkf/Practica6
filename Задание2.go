package main
import (
	"fmt"
	"sync"
)

type Cache struct {
	mu sync.RWMutex
	m  map[string]string
}
func (c *Cache) Get(k string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.m[k]
	return v, ok
}
func (c *Cache) Set(k, v string) {
	c.mu.Lock() 
	defer c.mu.Unlock()
	c.m[k] = v
}
func main() {
	c := Cache{m: make(map[string]string)}
	c.Set("a", "1")
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			v, ok := c.Get("a")
			fmt.Println("читатель", n, v, ok)
		}(i)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.Set("a", "2")
		fmt.Println("писатель обновил кэш")
	}()
	wg.Wait()
}