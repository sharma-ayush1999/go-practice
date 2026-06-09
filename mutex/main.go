package main

import (
	"fmt"
	"sync"
	"time"
)

// ── Thread-safe cache using RWMutex ───────────────────────────────────────────
// RWMutex allows many concurrent readers OR one exclusive writer.
// Use when reads >> writes (typical for a cache).
type Cache struct {
	mu 		sync.RWMutex // protects the map
	store 	map[string]string
}

func NewCache() *Cache { return &Cache{store: make(map[string]string)} }

func (c *Cache) Get (key string) (string, bool) {
	c.mu.RLock()	// acquire READ lock — multiple goroutines can hold this simultaneously
	defer c.mu.RUnlock()

	v, ok := c.store[key]
	return v, ok
}

func (c *Cache) Set(key, val string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.store[key] = val
}


func (c *Cache) Delete(key string){
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.store, key)
}

// ── Thread-safe counter using sync.Mutex ──────────────────────────────────────
// Mutex (not RWMutex) because every operation is a write.

type Counter struct {
	mu		sync.Mutex
	value 	int
}

func (c *Counter) Increment(){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}


func main(){
	cache := NewCache()
	counter := &Counter{}
	var wg sync.WaitGroup

	// Launch 50 goroutines: 40 readers, 10 writers — all concurrent.
	for i := range 50 {
		wg.Add(1)
		go func(id int){
			defer wg.Done()
			if id % 5 == 0 {
				// Writer goroutine
				key := fmt.Sprintf("key-%d", id)
				cache.Set(key, fmt.Sprintf("val-%d", id))
				counter.Increment()
			} else {
				// Reader goroutine — concurrent with other readers
				cache.Get(fmt.Sprintf("key-%d", id-id%5))
			}
			time.Sleep(1 * time.Millisecond)
		}(i)
	}

	wg.Wait()
	fmt.Printf("writes: %d\n", counter.value)
}