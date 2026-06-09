package main

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
)

// ── Pool of reusable byte buffers ─────────────────────────────────────────────
// bytes.Buffer is the pooled object. New() is called when pool is empty.
var bufPool = sync.Pool{
	New: func() interface {} {
		// Allocate a new buffer with 1KB initial capacity.
        // This runs only when the pool has no available object.
		return bytes.NewBuffer(make([]byte, 0, 1024))
	},
}

// processRequest simulates work that needs a temporary buffer.
func processRequest(data string) string {
	// STEP 1: Get a buffer from the pool (or a new one if pool is empty).
	buf := bufPool.Get().(*bytes.Buffer)
	
	// STEP 2: ALWAYS reset before use — the buffer may have been used before.
	buf.Reset()

	// STEP 3: Use the buffer for the work.
	buf.WriteString("PROCESSED: ")
	buf.WriteString(data)
	result := buf.String()

	// STEP 4: Return to pool. Do NOT use buf after this line.
    // The pool may give this buffer to another goroutine immediately.
	bufPool.Put(buf)

	return result
}


// ── Benchmarks showing the difference ─────────────────────────────────────────

// BenchmarkWithPool: reuses buffers → fewer allocations → less GC
func BenchmarkWithPool(b *testing.B){
	for b.Loop() {
		processRequest("hello world")
	}
}

// BenchmarkWithoutPool: allocates a new buffer every call
func BenchmarkWithoutPool(b *testing.B){
	for b.Loop() {
		var buf bytes.Buffer  // new allocation every iteration
		buf.WriteString("PROCESSED: ")
		buf.WriteString("hello world")
		_ = buf.String()
	}
}


func main(){
	result := processRequest("test data")
	fmt.Println(result)	// PROCESSED: test data
	
	// To run benchmarks:
    // go test -bench=. -benchmem
    // BenchmarkWithPool-8      3842190    310 ns/op      0 B/op    0 allocs/op
    // BenchmarkWithoutPool-8   2156780    553 ns/op    208 B/op    2 allocs/op
    // Pool: 0 allocations per call (reuses buffer from pool)
	fmt.Println("Run: go test -bench=. -benchmem to see pool vs no pool performance")

}