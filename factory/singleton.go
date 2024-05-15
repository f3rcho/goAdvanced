package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

func (Database) CreateSingleConnection() {
	fmt.Println("Creating singleton for DB")
	time.Sleep(2 * time.Second)
	fmt.Println("Creation Done")
}

var db *Database
var lock sync.Mutex
var once sync.Once

// an alternative of lock is once, that just execute the func once
// func getDataBaseInstance() *Database {
// 	if db == nil {
// 		once.Do(
// 			func() {
// 				fmt.Println("Creating DB Connection")
// 				db = &Database{}
// 				db.CreateSingleConnection()
// 			})
// 	} else {
// 		fmt.Println("DB Already created")
// 	}
// 	return db
// }

func getDataBaseInstance() *Database {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		fmt.Println("Creating DB Connection")
		db = &Database{}
		db.CreateSingleConnection()
	} else {
		fmt.Println("DB Already created")
	}
	return db
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDataBaseInstance()
		}()
	}
	wg.Wait()
}
