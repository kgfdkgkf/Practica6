package main
import (
	"fmt"
	"sync"
)
type Queue struct {
	mu sync.Mutex
	q  []string
}
func (s *Queue) Push(x string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.q = append(s.q, x)
}
func (s *Queue) Pop() (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.q) == 0 {
		return "", false
	}
	x := s.q[0]
	s.q = s.q[1:]
	return x, true
}

func main() {
	var s Queue
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			s.Push(fmt.Sprint("задача ", n))
		}(i)
	}
	wg.Wait()
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			for {
				x, ok := s.Pop()
				if !ok {
					return
				}
				fmt.Println("обработчик", n, "взял:", x)
			}
		}(i)
	}
	wg.Wait()
}