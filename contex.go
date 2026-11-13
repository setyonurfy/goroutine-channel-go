package goroutine_go

import (
	"context"
	"time"
)

func KerjaLama(ctx context.Context, durasi time.Duration) error {
	select {
	case <-time.After(durasi):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func Pencacah(ctx context.Context) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := 1; ; i++ {
			select {
			case out <- 1:
			case <-ctx.Done():
				return

			}
		}
	}()
	return out
}
