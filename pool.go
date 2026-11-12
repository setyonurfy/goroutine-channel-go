package goroutine_go

import (
	"sync"
	"time"
)

type tugas struct {
	indeks int
	nilai  int
}

// WorkerPool menghitung kuadrat tiap angka memakai sejumlah pekerja.

func WorkerPool(jumlahPekerja int, angka []int) []int {
	jobs := make(chan tugas)
	semua := make([]int, len(angka))

	var wg sync.WaitGroup
	for w := 0; w < jumlahPekerja; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range jobs {
				time.Sleep(100 * time.Millisecond) // pura-pura kerja berat
				semua[t.indeks] = t.nilai * t.nilai
			}
		}()
	}

	for i, n := range angka {
		jobs <- tugas{indeks: i, nilai: n}
	}
	close(jobs)

	wg.Wait()
	return semua
}
