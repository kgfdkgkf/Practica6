package main

import (
	"fmt"
	"sync"
)
func main() {
	jobs := make(chan int)
	var mu sync.Mutex  
	var wgW, wgP sync.WaitGroup 
	sum := 0 
	count := 0 
	for w := 1; w <= 3; w++ {
		wgW.Add(1)
		go func(id int) {
			defer wgW.Done()
			for j := range jobs {
				mu.Lock()
				sum += j * j
				count++
				mu.Unlock()
			}
		}(w)
	}
	for p := 0; p < 2; p++ {
		wgP.Add(1)
		go func(p int) {
			defer wgP.Done()
			for i := 1; i <= 5; i++ {
				jobs <- p*10 + i
			}
		}(p)
	}
	wgP.Wait() 
	close(jobs) 
	wgW.Wait()
	fmt.Println("обработано задач:", count)
	fmt.Println("сумма квадратов:", sum)
}