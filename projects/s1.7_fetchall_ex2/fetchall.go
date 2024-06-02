package main

import (
	"context"
	"fmt"
	"time"
)

// A function that simulates some work, checks for a cancellation signal, and writes to a channel
func someFunction(ctx context.Context, msgChan chan<- string) {
	for {
		select {
		case <-ctx.Done():
			msgChan <- "Goroutine received cancellation signal"
			close(msgChan)
			return
		default:
			// Simulate some work
			msgChan <- "Goroutine working..."
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Ensure the context is cancelled when main returns

	// Create a channel for messages
	msgChan := make(chan string)

	// Start the goroutine
	go someFunction(ctx, msgChan)

	// Read and print messages from the channel
	for msg := range msgChan {
		fmt.Println(msg)
	}

	fmt.Println("Main function completed")
}
