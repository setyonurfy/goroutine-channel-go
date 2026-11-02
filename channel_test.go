package goroutine_go

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // Close the channel to signal that no more values will be sent
	go func() {
		time.Sleep(2 * time.Second)        // Simulate some work
		channel <- "Hello from goroutine!" // Send a value to the channel
		fmt.Println("Goroutine finished sending value")
	}()
	reciveChan := <-channel // Receive a value from the channel
	fmt.Println(reciveChan)
	time.Sleep(5 * time.Second) // Wait for the goroutine to finish

}
