package main

import (
	"fmt"
	"sync"
)
type Votes struct {
	mu sync.Mutex
	m  map[string]int
}
func (v *Votes) Add(name string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.m[name]++
}
func main() {
	v := Votes{m: make(map[string]int)}
	names := []string{"Иванов", "Петров", "Сидоров"}
	var wg sync.WaitGroup
	for i := 0; i < 99; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			v.Add(names[n%3])
		}(i)
	}
	wg.Wait()
	v.mu.Lock()
	for k, x := range v.m {
		fmt.Println(k, x)
	}
	v.mu.Unlock()
}