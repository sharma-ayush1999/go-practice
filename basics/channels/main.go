package main

import "fmt"

func producer(ch chan<- int){
	for i := range 5{
		ch <- i	// send — blocks if buffer is full
	}
	close(ch)	// signal: no more values	
}

func main(){
	// ── Unbuffered channel ──────────────────────────────
	unbufd := make(chan string)
	go func(){
		unbufd <- "ping"  // blocks until main reads
	}()

	msg := <-unbufd	// blocks until goroutine sends
	fmt.Println("unbuffered channel", msg)

	// ── Buffered channel ────────────────────────────────
    // capacity 3: first 3 sends don't block
	bufd := make(chan int, 3)
	bufd <- 1
	bufd <- 2
	bufd <- 3
	bufd <- 4 // would block — buffer full

	fmt.Println(<-bufd) //1
	fmt.Println(<-bufd) //2

    // ── range over channel ──────────────────────────────
	jobs := make(chan int, 5)
	go producer(jobs)

	// range reads until the channel is closed
	for n := range jobs {
		fmt.Println("received:", n)
	}

    // ── Two-value receive (detect close) ────────────────
	ch := make(chan int, 2)
	ch <- 10
	close(ch)

	v1, ok := <-ch // 10, true
	v2, ok2 := <-ch	// 0, false (channel closed)
	fmt.Println(v1, ok)  // 10 true
	fmt.Println(v2, ok2) // 0 false

}