package main

import (
	"fmt"
)

// Package-level: must use var (not :=)
var appName string = "myApp"
var version = 1  // type inferred as int

const Pi = 3.14159 // constant — immutable


func main() {
	// Inside a function: := is the idiomatic shorthand
	name := "Ayush" // string, inferred
	age := 27 // int, inferred
	var score float64 = 98.5 // explicit type

	// Zero values — Go always initializes variables
	var count int //0
	var flag bool //false
	var msg string // ""

	fmt.Println(name, age, score)
	fmt.Println("zeros:", count, flag, msg)
	fmt.Println(appName, version, Pi)

}