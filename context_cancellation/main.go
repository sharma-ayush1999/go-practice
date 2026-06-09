package main

import (
	"context"
	"fmt"
	"time"
)

// worker does periodic work and stops when ctx is cancelled.
// ctx.Done() returns a channel that is closed when the context is cancelled.
func worker(ctx context.Context, id int){
	for{
		select {
		case <-ctx.Done():
			// ctx.Err() tells us WHY: context.Canceled or context.DeadlineExceeded
			fmt.Printf("worker %d stopping: %v\n", id, ctx.Err())
			return // exit the goroutine cleanly
		case <-time.After(300 * time.Microsecond):
			// Do one unit of work. In production: poll queue, process event, etc.
			fmt.Printf("worker %d working...\n", id)
		}
	}
}


func main(){
	// WithTimeout cancels the context automatically after 1 second.
    // WithCancel would give manual control via cancel().
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)

	// ALWAYS call cancel() — releases resources even if timeout fires first.
    // defer ensures it runs even if main panics.
	defer cancel()

	// Start 3 workers — all share the same context.
	for i:= range 3 {
		go worker(ctx, i)
	}

    // Wait for context to expire (1 second).
	<-ctx.Done()

	// Give workers a moment to print their stopping message.
	time.Sleep(100 * time.Millisecond)
	fmt.Println("all workers stopped")
}