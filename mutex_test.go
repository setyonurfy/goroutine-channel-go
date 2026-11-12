package goroutine_go

import (
	"sync"
	"testing"
)

func TestCounterTanpaMutex(t *testing.T) {
	counter := 0

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}
	wg.Wait()

	if counter != 1000 {
		t.Errorf("got %d, must be 1000", counter)
	}
}

func TestSafeCounter(t *testing.T) {
	var c SafeCounter
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

func TestSafeCounter_BacaTulisBersamaan(t *testing.T) {
	var c SafeCounter
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			c.Tambah()
		}()
		go func() {
			defer wg.Done()
			_ = c.Nilai()
		}()
	}
	wg.Wait()

	if got := c.nilai; got != 100 {
		t.Errorf("got %d, must be 100", got)
	}
}
