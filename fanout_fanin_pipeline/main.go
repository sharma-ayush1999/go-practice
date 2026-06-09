package main

import (
	"fmt"
	"sync"
)

// generate sends numbers 1..n into a channel, then closes it.
// Closing signals all receivers that no more data is coming.
func generate(n int) <-chan int {
	out := make(chan int) // unbuffered — generator blocks until consumer reads

	go func(){
		for i := 1; i <= n; i++ {
			out <- i // send each number
		}
		close(out) // IMPORTANT: close after all sends — receivers exit their range loops
	}()
	return out   // return the read-only end of the channel
}


// square reads from 'in' and sends each value squared to 'out'.
// This is one worker — we'll run several in parallel.
func square(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // signal WaitGroup when this worker exits
	for n:= range in {
		// 'range' over a channel reads until the channel is closed.
		out <- n * n // send squared value to shared output channel
	}
}


// merge collects from the output channel and closes it when all workers finish.
func merge(out chan int, wg *sync.WaitGroup){
	// Wait in a goroutine so merge() returns immediately.
	go func(){
		wg.Wait() // block until all workers call Done()
		close(out) // close the output channel — signals the printer to stop
	}()
}


func main(){
	// STEP 1: Create pipeline source (generates 1..10).
	nums := generate(10)

	// STEP 2: Create shared output channel (buffered to prevent workers blocking).
	results := make(chan int, 10)

	// STEP 3: Start 3 worker goroutines (fan-out).
    // All workers read from the SAME 'nums' channel — Go's channel guarantees
    // each value is received by exactly one goroutine (safe without mutex).
	var wg sync.WaitGroup
	for i := range 3{
		wg.Add(i)
		go square(nums, results, &wg)
	}

	// STEP 4: Close 'results' after all workers finish (fan-in).
	merge(results, &wg)

	// STEP 5: Print all results. Loop exits when 'results' is closed.
	for r := range results {
		fmt.Println(r)
	}
}

// Output (order varies because workers run concurrently):
// 1, 4, 9, 16, 25, 36, 49, 64, 81, 100
