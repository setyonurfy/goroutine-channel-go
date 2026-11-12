package goroutine_go

import (
	"reflect"

	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	got := WorkerPool(3, []int{1, 2, 3, 4, 5})

	want := []int{1, 4, 9, 16, 25}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("dapat %v, harusnya %v", got, want)
	}
}

func TestWorkerPool_LebihCepat(t *testing.T) {
	mulai := time.Now()
	WorkerPool(3, []int{1, 2, 3, 4, 5, 6})
	durasi := time.Since(mulai)

	if durasi > 4*time.Second {
		t.Errorf("terlalu lama (%v), pekerja tidak bekerja bersama", durasi)
	}
}

func TestWorkerPool_HasilBenar(t *testing.T) {
	got := WorkerPool(3, []int{1, 2, 3, 4, 5})

	want := []int{1, 4, 9, 16, 25}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("dapat %v, harusnya %v", got, want)
	}
}
