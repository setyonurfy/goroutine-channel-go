package goroutine_go

import (
	"fmt"
	"testing"
	"time"
)

func TestUnduh(t *testing.T) {
	got := Unduh(1)
	want := "File ke-1 selesai"

	if got != want {
		t.Errorf("dapat %q, harusnya %q", got, want)

	}
}

func TestUnduhSemua_HasilBenar(t *testing.T) {
	hasil := UnduhSemua(5)
	if len(hasil) != 5 {
		t.Fatalf("jumlah hasil %d, harusnya 5", len(hasil))
	}
	for i, pesan := range hasil {
		want := fmt.Sprintf("File ke-%d selesai", i+1)
		if pesan != want {
			t.Errorf("hasil[%d] = %q, harusnya %q", i, pesan, want)
		}
	}
}

func TestUnduhSemua_Bersamaan(t *testing.T) {
	mulai := time.Now()
	UnduhSemua(5)
	durasi := time.Since(mulai)

	if durasi > 1500*time.Millisecond {
		t.Errorf("terlalu lama (%v), sepertinya tidak berjalan bersamaan", durasi)
	}
}
