package goroutine_go

import (
	"fmt"
	"testing"
	"time"
)

func RunGoroutine() {
	fmt.Println("goroutine 1")
}

func TestGoroutine(t *testing.T) {
	go RunGoroutine()
	fmt.Println("goroutine 2")

	time.Sleep(1 * time.Second) // Wait for the goroutine to finish
}

func DisplayNumber(number int) {
	fmt.Println("display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(5 * time.Second) // Wait for the goroutines to finish
}
