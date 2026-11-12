package goroutine_go

import (
	"errors"
	"testing"
	"time"
)

func kirimSetelah(jeda time.Duration, pesan string) <-chan string {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(jeda)
		ch <- pesan
	}()
	return ch
}

func TestTercepat(t *testing.T) {
	lambat := kirimSetelah(3*time.Second, "lambat")
	cepat := kirimSetelah(1*time.Second, "cepat")

	got := Tercepat(lambat, cepat)

	if got != "cepat" {
		t.Errorf("got %q, must be %q", got, "cepat")
	}
}

func TestTungguDenganTimeout_Sukses(t *testing.T) {
	ch := kirimSetelah(1*time.Second, "ok")

	got, err := TungguDenganTimeout(ch, 5*time.Second)
	if err != nil {
		t.Fatalf("error tidak diharapkan: %v", err)
	}
	if got != "ok" {
		t.Errorf("got %q, must be %q", got, "ok")
	}
}

func TestTungguDenganTimeout_Habis(t *testing.T) {
	ch := kirimSetelah(5*time.Second, "terlambat")

	_, err := TungguDenganTimeout(ch, 1*time.Second)
	if !errors.Is(err, ErrTimeout) {
		t.Errorf("got error %v, must be ErrTimeout", err)
	}
}

func TestCobaAmbil(t *testing.T) {
	ch := make(chan int, 1)

	if _, ok := CobaAmbil(ch); ok {
		t.Error("chanel kosong, harusnya ok = false")
	}

	ch <- 42
	n, ok := CobaAmbil(ch)
	if !ok || n != 42 {
		t.Errorf("got (%d, %v), must be (42, true)", n, ok)
	}
}

func TestTercepatDenganTimeout_Sukses(t *testing.T) {
	a := kirimSetelah(3*time.Second, "lambat")
	b := kirimSetelah(1*time.Second, "cepat")

	got, err := TercepatDenganTimeout(a, b, 5*time.Second)
	if err != nil {
		t.Fatalf("error tidak diharapkan : %v", err)
	}
	if got != "cepat" {
		t.Errorf("got %q, must be %q", got, "cepat")
	}
}

func TestTercepatDenganTimeout_Habis(t *testing.T) {
	a := kirimSetelah(5*time.Second, "a")
	b := kirimSetelah(3*time.Second, "b")

	_, err := TercepatDenganTimeout(a, b, 1*time.Second)
	if !errors.Is(err, ErrTimeout) {
		t.Errorf("got error %v, must be ErrTimeout", err)
	}
}
