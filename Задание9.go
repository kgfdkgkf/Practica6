package main

import (
	"fmt"
	"sync"
)
type Metrics struct {
	mu    sync.Mutex
	ok    int
	fail  int
	total int 
}
func (m *Metrics) Add(success bool, ms int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if success {
		m.ok++
	} else {
		m.fail++
	}
	m.total += ms
}
func (m *Metrics) Report() {
	m.mu.Lock()
	defer m.mu.Unlock()
	n := m.ok + m.fail
	avg := 0
	if n > 0 { 
		avg = m.total / n
	}
	fmt.Println("===== ОТЧЁТ =====")
	fmt.Println("успешных запросов:  ", m.ok)
	fmt.Println("неуспешных запросов:", m.fail)
	fmt.Println("всего запросов:     ", n)
	fmt.Println("среднее время, мс:  ", avg)
	fmt.Println("=================")
}

func main() {
	var m Metrics
	var wg sync.WaitGroup
	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			m.Add(n%5 != 0, n*3)
		}(i)
	}
	wg.Wait()
	m.Report()
}