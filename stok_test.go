package goroutine_go

import (
	"errors"
	"sync"
	"testing"
)

func TestSafeStok_TambahBersamaan(t *testing.T) {
	s := NewSafeStok()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.Tambah("laptop", 1)
		}()
	}
	wg.Wait()

	if got := s.Cek("laptop"); got != 100 {
		t.Errorf("got %d, must be 100", got)
	}
}

func TestSafeStok_KurangiTidakCukup(t *testing.T) {
	s := NewSafeStok()
	s.Tambah("laptop", 5)

	err := s.Kurangi("laptop", 10)
	if !errors.Is(err, ErrStokKurang) {
		t.Errorf("got error %v, must be ErrStokKurang", err)
	}
	if got := s.Cek("laptop"); got != 5 {
		t.Errorf("stok jadi %d, harusnya tetap 5", got)
	}
}

func TestSafeStok_KurangiBersamaan(t *testing.T) {
	s := NewSafeStok()
	s.Tambah("laptop", 10)

	var berhasil SafeCounter
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := s.Kurangi("laptop", 1); err == nil {
				berhasil.Tambah()
			}
		}()
	}
	wg.Wait()

	if got := berhasil.Nilai(); got != 10 {
		t.Errorf("yang berhasil %d, harusnya 10", got)
	}
	if got := s.Cek("laptop"); got != 0 {
		t.Errorf("stok akhir %d, harusnya 0", got)
	}
}
