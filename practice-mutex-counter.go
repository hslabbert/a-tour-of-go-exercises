package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

// LoopInc increments the value of the counter for the given key
// in a loop of count cycles.
func (c *SafeCounter) LoopInc(key string, count int) {
	for i := 0; i < count; i++ {
		go c.Inc(key)
	}
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	go c.LoopInc("somekey", 1000)
	go c.LoopInc("somekey", 143)
	go c.LoopInc("anotherkey", 100)
	go c.LoopInc("somekey", 547)
	go c.LoopInc("anotherkey", 300)

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
	fmt.Println(c.Value("anotherkey"))
}
