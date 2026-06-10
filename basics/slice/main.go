package main

import (
	"fmt"
)


func main(){
	 // Array — fixed size, rarely used directly
	 arr := [3]int{1,2,3}
	 fmt.Println("array:", arr)

	// Slice — dynamic, the idiomatic choice
	s := []int{1,2,3}
	fmt.Println("Slice:", s, "len:", len(s), "cap:", cap(s))

	// append — grows the slice
	s = append(s, 40, 50)
	fmt.Println("after append:", s)

	// make([]T, len, cap) — pre-allocate when you know the size
    // avoids repeated re-allocation inside a loop
	pre := make([]int, 0, 100)
	for i := range 5 {
		pre = append(pre, i*i)
	}
	fmt.Println("pre-allocated:", pre)

	// Slicing — creates a view, shares memory
	view := s[1:3]	// elements at index 1 and 2 (not 3)
	fmt.Println("view:", view)
	view[0] = 999	// modifies the original slice too!
	fmt.Println("modified:", view)

	// Copy — independent copy, no shared memory
	dst := make([]int, len(s))
	copy(dst, s)
	dst[0] = 0
	fmt.Println("src[0]:", s[0], "dst[0]:", dst[0])

}