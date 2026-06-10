package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchData(id int, wg *sync.WaitGroup) {
	defer wg.Done()	// decrement counter when this goroutine exits
	
	// simulate network call
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("goroutine %d: data fetched\n", id)
}

func main(){
	var wg sync.WaitGroup

	start := time.Now()

	// Launch 5 goroutines concurrently
	for i := range 5{
		wg.Add(1)	// increment counter before launching
		go fetchData(i, &wg)
	}

	wg.Wait() // block until all goroutines call wg.Done()

	elapsed := time.Since(start)
	// All 5 fetches ran in ~100ms (concurrent), not 500ms (sequential)
	fmt.Printf("all done in %v\n", elapsed)
}