package main

import (
	"fmt"
	"sync"
)
type Cinema struct {
	mu    sync.Mutex
	seats []bool
}
func (c *Cinema) Book(n int) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if n < 1 || n > 38 {
		return false
	}
	if c.seats[n-1] {
		return false
	}
	c.seats[n-1] = true
	return true
}

func main() {
	c := Cinema{seats: make([]bool, 38)}
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(u int) {
			defer wg.Done()
			if c.Book(7) {
				fmt.Println("пользователь", u, "забронировал место 7")
			} else {
				fmt.Println("пользователь", u, "не успел, место 7 занято")
			}
		}(i)
	}
	wg.Wait()
}