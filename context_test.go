package goroutine_go

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestKerjaLama_SelesaiNormal(t *testing.T) {
	err := KerjaLama(context.Background(), 50*time.Millisecond)
	if err != nil {
		t.Errorf("error tidak duharapkan: %v", err)
	}
}

func TestKerjaLama_Dibatalkan(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()
	mulai := time.Now()
	err := KerjaLama(ctx, 5*time.Second)

	if !errors.Is(err, context.Canceled) {
		t.Errorf("got error %v, must be context.Canceled", err)
	}
	if durasi := time.Since(mulai); durasi > time.Second {
		t.Errorf("berhenti terlalu lama (%v), harusnya segera setelah cancel", durasi)
	}
}

func TestKerjaLama_timeout(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := KerjaLama(ctx, 5*time.Second)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("got error %v, must be context.DeadlineExceeded", err)
	}
}

func TestPencacah_BerhentiSetelahBatal(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := Pencacah(ctx)

	for i := 1; i <= 3; i++ {
		if got := <-ch; got != i {
			t.Errorf("got %d, must be %d", got, i)
		}
	}
	cancel()

	// channel harus tertutup

	batas := time.After(time.Second)
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return
			}
		case <-batas:
			t.Fatal("cahnnel tidak tertutup setelah cancel: goroutine bocor")
		}
	}
}
