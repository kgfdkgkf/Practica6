package main

import (
	"fmt"
	"sync"
)
type Logger struct {
	mu sync.Mutex
}
func (l *Logger) Log(id int, s string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Printf("[горутина %d] %s\n", id, s)
}
func main() {
	var l Logger
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			l.Log(n, "начала работу")
			l.Log(n, "работает")
			l.Log(n, "закончила работу")
		}(i)
	}
	wg.Wait()
}