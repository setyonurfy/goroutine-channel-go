package goroutine_go

import (
	"errors"
	"time"
)

var ErrTimeout = errors.New("waktu habis")

func Tercepat(a, b <-chan string) string {
	select {
	case pesan := <-a:
		return pesan
	case pesan := <-b:
		return pesan
	}
}

func TungguDenganTimeout(ch <-chan string, batas time.Duration) (string, error) {
	select {
	case pesan := <-ch:
		return pesan, nil
	case <-time.After(batas):
		return "", ErrTimeout
	}
}

func CobaAmbil(ch <-chan int) (int, bool) {
	select {
	case n := <-ch:
		return n, true
	default:
		return 0, false
	}
}

func TercepatDenganTimeout(a, b <-chan string, batas time.Duration) (string, error) {
	select {
	case pesan := <-a:
		return pesan, nil
	case pesan := <-b:
		return pesan, nil
	case <-time.After(batas):
		return "", ErrTimeout
	}
}
