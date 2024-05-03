package sync

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Deposit(amount int, wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	lock.Lock()
	b := balance
	balance = b + amount
	lock.Unlock()
}

func Balance() int {
	b := balance
	return b
}

func RunDeposit() {
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}
	for i := 1; i <= 5; i++ {
		wg.Add(3)
		go Deposit(i*100, &wg, &lock)
		go Deposit(i*200, &wg, &lock)
		go Deposit(i*300, &wg, &lock)
	}
	wg.Wait()
	fmt.Println(Balance())
}
