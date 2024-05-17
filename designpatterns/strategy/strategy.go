package main

import (
	"fmt"
)

// strategy interface

type EvictionAlgo interface {
	evict(c *Cache)
}

// concrete strategy

type Fifo struct{}

// Primera en entrar, primera en salir
func (f *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by FIFO strategy")
}

// concrete strategy

type Lru struct{}

// Menos usada recientemente
func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by lru strategy")
}

type Lfu struct{}

// Menos frecuentemente usada (LFU)
func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by Lfu strategy")
}

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     2,
		maxCapacity:  3,
	}
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) add(key, value string) {
	fmt.Printf("adding key:%s, with value: %s\n", key, value)
	if c.capacity == c.maxCapacity {
		fmt.Println("capacity full")
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get(key string) {
	fmt.Printf("deleting key:%s\n", key)
	delete(c.storage, key)
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

func main() {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.add("aaa", "1")
	cache.add("bbb", "2")
	cache.add("ccc", "3")

	lru := &Lfu{}
	cache.setEvictionAlgo(lru)

	cache.add("ddd", "4")

	fifo := &Fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("eee", "5")

	cache.get("aaa")
}
