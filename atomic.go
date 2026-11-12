package goroutine_go

import "sync/atomic"

type AtomicCounter struct {
	nilai atomic.Int64
}

func (c *AtomicCounter) Tambah() {
	c.nilai.Add(1)
}
func (c *AtomicCounter) Nilai() int64 {
	return c.nilai.Load()
}
