# ADVANCED TOPICS


## RACE CONDITIONS

In go, when we have many go routines accessing to same source of data we could face a race condition. A way to resolve this issue it's to use the sync module. Locking the execution to ensure that one routine at a tiem can access it. ```lock *sync.Mutex```

```go
func Deposit(amount int, wg *sync.WaitGroup, lock *sync.Mutex) {
  defer wg.Done()
	lock.Lock()
	b := balance
	balance = b + amount
	lock.Unlock()
}
```
### Using Sync Mutex
- sync.Mutex.Lock() help us to lock the access to shared variables between goroutines.
- sync.Mutex.Unlock() will unlooked the value needed to access.

### How to know if there is a RACE condition
```bash
go build --race
```
then execute the binary file.

### RWMutex

Read and write mutex, allow us to lock writes and reads. In this case we lock the writes to 1 at a time, and the reads N at a time using RLock() and RUnlock()
1 -> write
N -> read

### Cache system