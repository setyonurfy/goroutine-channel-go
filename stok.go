package goroutine_go

import (
	"errors"
	"sync"
)

var ErrStokKurang = errors.New("stok kurang")

type SafeStok struct {
	mu   sync.Mutex
	data map[string]int
}

func NewSafeStok() *SafeStok {
	return &SafeStok{data: make(map[string]int)}
}

func (s *SafeStok) Tambah(nama string, jumlah int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[nama] += jumlah
}

func (s *SafeStok) Cek(nama string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.data[nama]
}

func (s *SafeStok) Kurangi(nama string, jumlah int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.data[nama] < jumlah {
		return ErrStokKurang
	}
	s.data[nama] -= jumlah
	return nil
}
