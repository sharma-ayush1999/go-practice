package main

import (
	"fmt"
	"sync"
)

func makeCounter() func() int {
	count := 0	// captured by the closure
	return func () int {
		count++	// count persists between calls
		return count
	}
}

func main(){
	counter := makeCounter()
	fmt.Println(counter()) //1
	fmt.Println(counter()) //2
	fmt.Println(counter()) //3

	counter2 := makeCounter() //independent counter
	fmt.Println(counter2()) //1

	// ── PITFALL: loop variable capture ───────────────────────
	var wg sync.WaitGroup
	results := make([]int, 5)

	// WRONG — all goroutines see the same 'i' variable
    // By the time they run, i == 5
	for i := range 5 {
		wg.Add(1)
		go func(){
			defer wg.Done()
			results[i] = i  // DATA RACE + wrong value
		}()
	}
	// Don't actually run this — it's broken on purpose


	wg.Wait()
	fmt.Println("WRONG answer", results)

	// FIX 1: pass i as an argument
	results2 := make([]int, 5)
	var wg2 sync.WaitGroup
	for i := range 5 {
		wg2.Add(1)
		go func (n int) { // n is a copy of i at this moment
			defer wg2.Done()
			results2[n] = n * n
		}(i)
	}

	wg2.Wait()
	fmt.Println("squares:", results2)
}	