package main

import (
	"fmt"
)

func main(){
    // Declare and initialize with a literal
	sources := map[string]int{
		"Alice": 90,
		"Bob": 85,
	}

	// Add / update a key
	sources["charlie"] = 54
	sources["Alice"] = 34

	fmt.Println("Alice:", sources["Alice"])

	// Two-value lookup — always use this for safety
	val, ok := sources["Dave"]
	if !ok {
		fmt.Println("Key not found, zero value", val)
	}

	// Delete a key
	delete(sources, "Alice")

	// Iterate — order is random every run (by design)
	for name, score := range sources {
		fmt.Printf("%s: %d\n", name, score)
	}

	// make() — empty map ready to use
	inventory := make(map[string]int)
	inventory["apples"] = 65
	fmt.Println("Inventory", inventory)
}
