package goroutine_go

import "sync"

type SafeCounter struct {
	mu    sync.Mutex
	nilai int
}

func (c *SafeCounter) Tambah() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.nilai++
}

func (c *SafeCounter) Nilai() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.nilai
}
