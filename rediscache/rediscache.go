package rediscache

import (
	"fmt"
	"sync"

	"github.com/f3rcho/goAdvanced/cache"
)

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}
type Service struct {
	InProgress map[int]bool
	IsPending  map[int][]chan FunctionResult
	Lock       sync.RWMutex
	f          Function
	cache      map[int]FunctionResult
}

func (s *Service) Work(job int) FunctionResult {
	s.Lock.RLock() //reads blocks
	exists := s.InProgress[job]
	if exists {
		s.Lock.RUnlock() // unblock
		response := make(chan FunctionResult)
		defer close(response)

		s.Lock.Lock()
		s.IsPending[job] = append(s.IsPending[job], response)
		s.Lock.Unlock()

		fmt.Printf("Waiting for response job:%d\n", job)

		return <-response
	}
	s.Lock.RUnlock()

	s.Lock.RLock()
	result, exists := s.cache[job]
	s.Lock.RUnlock()

	if !exists {
		s.Lock.Lock()
		s.InProgress[job] = true
		s.Lock.Unlock()

		fmt.Printf("Calculate fib for %d\n", job)
		result.value, result.err = s.f(job)

		s.Lock.RLock()
		pendingWorkers, exists := s.IsPending[job]
		s.Lock.RUnlock()

		if exists {
			for _, pendingW := range pendingWorkers {
				pendingW <- result
			}
		}

		s.Lock.Lock()
		s.InProgress[job] = false
		s.IsPending[job] = make([]chan FunctionResult, 0)
		s.cache[job] = result
		s.Lock.Unlock()
	}
	return result
}

func NewService(f Function) *Service {
	return &Service{
		InProgress: make(map[int]bool),
		IsPending:  make(map[int][]chan FunctionResult),
		f:          f,
		cache:      make(map[int]FunctionResult),
	}
}

func RedisCache() {
	service := NewService(cache.GetFibonacci)
	jobs := []int{5, 5, 6, 7, 8}
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
