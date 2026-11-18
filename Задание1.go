package main
import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup 
	c := 0 

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done() 
			mu.Lock()       
			c++             
			mu.Unlock() 
		}()
	}

	wg.Wait() // ждём все 100 горутин
	fmt.Println("просмотров:", c)
}