package goroutine_go

import (
	"sync"
	"sync/atomic"
	"testing"
)

func Test(t *testing.T) {
	var c AtomicCounter
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Tambah()
		}()
	}
	wg.Wait()

	if got := c.Nilai(); got != 1000 {
		t.Errorf("got %d, must be 1000", got)
	}
}

func TestCompareAndSwap_HanyaSatuPemenang(t *testing.T) {
	var status atomic.Int64
	var pemenang AtomicCounter
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if status.CompareAndSwap(0, 1) {
				pemenang.Tambah()
			}
		}()
	}
	wg.Wait()

	if got := pemenang.Nilai(); got != 1 {
		t.Errorf("got %d,  must be 1", got)
	}
}
