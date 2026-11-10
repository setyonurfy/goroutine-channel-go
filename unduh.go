package goroutine_go

import (
	"fmt"
	"sync"
	"time"
)

func Unduh(id int) string {
	time.Sleep(time.Duration(id) * 2 * time.Second)
	return fmt.Sprintf("File ke-%d selesai", id)
}

func UnduhSemua(n int) []string {
	hasil := make([]string, n)
	var wg sync.WaitGroup

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			hasil[id-1] = Unduh(id)
		}(i)
	}
	wg.Wait()
	return hasil
}
