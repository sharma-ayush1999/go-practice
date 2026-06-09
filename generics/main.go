package main

import (
	"fmt"

	"golang.org/x/exp/constraints" // go get golang.org/x/exp
)

// ── Map: transform each element of a slice ─────────────────────────────────────
// [T, U any] means T and U can be ANY type.
// Without generics you'd need one function per type or use interface{}/reflect.
func Map[T, U any] (s []T, fn func(T) U) []U{
	result := make([]U, len(s))

	for i, v := range s{
		result[i] = fn(v)  // apply fn to each element
	}
	return result
}

// ── Filter: keep elements satisfying a predicate ───────────────────────────────
func Filter[T any] (s []T, fn func(T) bool) []T {
	var result[] T
	for _, v := range s{
		if fn(v) { result = append(result, v)}
	}
	return result
}

// ── Min: finds minimum value — constrained to ordered types ──────────────────
// constraints.Ordered includes all integers, floats, and strings.
func Min[T constraints.Ordered] (a, b T) T {
	if a < b { return a }
	return b
}

// ── Stack: generic LIFO data structure ──────────────────────────────────────
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push (v T) { s.items = append(s.items, v) }
func (s *Stack[T]) Len () int { return len(s.items) }
func (s *Stack[T]) Pop() (T, bool) {
	var zero T // zero value of T (0 for int, "" for string, nil for pointer)
	if len(s.items) == 0 { return zero, false }
	top := s.items[len(s.items) - 1]
	s.items = s.items[:len(s.items) - 1]
	return top, true
}

func main(){
    // Map: double each integer
	nums := []int{1,2,3,4,5}
	doubled := Map(nums, func (n int) int { return n * 2 } )
	fmt.Println(doubled)

	// Map: convert ints to strings
	strs := Map(nums, func (n int) string { return fmt.Sprintf("item-%d", n)})
	fmt.Println(strs)

	// Filter: keep only even numbers
	evens := Filter(nums, func (n int) bool { return n % 2 == 0})
	fmt.Println(evens)

	// Min: works for int and string with same function
	fmt.Println(Min(3, 7))
	fmt.Println(Min("b", "a"))

	
	// Generic stack
	var s Stack[string]
	s.Push("first")
	s.Push("second")
	top, _ := s.Pop()
	fmt.Println(top, s.Len())

}