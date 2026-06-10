package main

import (
	"context"
	"fmt"
	"time"
)

// Simulates a slow database query that respects context cancellation
func queryDb(ctx context.Context, query string) (string, error){
	done := make(chan string, 1)

	go func(){
		time.Sleep(200 * time.Millisecond) //simulate slow query
		done <- "result:" + query
	}()

	select {
	case result := <-done:
		return result, nil
	case <-ctx.Done():
		//ctx.Err() tells you why: DeadlineExceeded or Canceled
		return "", ctx.Err()
	}
}

func main(){
	// ── WithTimeout ─────────────────────────────────────────
	shortCtx, cancel2 := context.WithTimeout(context.Background(), 50 * time.Millisecond)
	defer cancel2()

	_, err2 := queryDb(shortCtx, "SELECT * FROM orders")
	fmt.Println("short timeout error:", err2) // context deadline exceeded

	// ── WithValue — pass request-scoped data ─────────────────
	type ctxKey string
	ctx3 := context.WithValue(context.Background(), ctxKey("requestID"), "req-123")

	
	// Retrieve value
	if rid, ok := ctx3.Value(ctxKey("requestID")).(string); ok {
		fmt.Println("request ID:", rid)
	}
}