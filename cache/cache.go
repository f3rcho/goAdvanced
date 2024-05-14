package cache

import (
	"fmt"
	"sync"
	"time"
)

func ExpensiveFib(n int) int {
	fmt.Printf("Calculate expensive fib for %d\n", n)
	time.Sleep(5 * time.Second)
	return n
}

type Service struct {
	InProgress map[int]bool
	IsPending  map[int][]chan int
	Lock       sync.RWMutex
}

func (s *Service) Work(job int) {
	s.Lock.RLock() //reads blocks
	exists := s.InProgress[job]
	if exists {
		s.Lock.RUnlock() // unblock
		response := make(chan int)
		defer close(response)
		s.Lock.Lock()
		s.IsPending[job] = append(s.IsPending[job], response)
		s.Lock.Unlock()
		fmt.Printf("Waiting for response job:%d\n", job)
		resp := <-response
		s.Lock.RLock() //reads blocks
		fmt.Printf("Response done, received %d\n", resp)
	}
	s.Lock.RUnlock()

	s.Lock.Lock()
	s.InProgress[job] = true
	s.Lock.Unlock()

	fmt.Printf("Calculate fib for %d\n", job)
	result := ExpensiveFib(job)

	s.Lock.RLock()
	pendingWorkers, exists := s.IsPending[job]
	s.Lock.RUnlock()
	if exists {
		for _, pendingW := range pendingWorkers {
			pendingW <- result
		}
		fmt.Printf("Result sent - all pending workers ready job:%d\n", job)
	}

	s.Lock.Lock()
	s.InProgress[job] = false
	s.IsPending[job] = make([]chan int, 0)
	s.Lock.Unlock()
}

func NewService() *Service {
	return &Service{
		InProgress: make(map[int]bool),
		IsPending:  make(map[int][]chan int),
	}
}

func RunExpensiveFib() {
	service := NewService()
	jobs := []int{3, 4, 5, 6, 7, 8, 3, 5}
	wg := sync.WaitGroup{}
	wg.Add(len(jobs))
	for _, n := range jobs {
		go func(job int) {
			defer wg.Done()
			service.Work(job)
		}(n)
	}
	wg.Wait()
}
