package main

import (
	"fmt"
	"math"
)

// Shape interface — any type with Area() and Perimeter() satisfies it
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct { Radius float64 }
type Rect struct { W, H float64}

// Circle satisfies Shape (no declaration needed)
func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

 
// Rect satisfies Shape
func (r Rect) Area() float64 { return r.H * r.W }
func (r Rect) Perimeter() float64 { return 2 * (r.H + r.W) }

// printShape accepts any Shape — polymorphism
func printShape(s Shape){
	fmt.Printf("Area=%.2f Perimeter=%.2f", s.Area(), s.Perimeter())
}

func main() {
	shapes := []Shape{
		Circle{Radius: 5},
		Rect{W: 5, H:4},
	}

	for _, s := range shapes {
		printShape(s)
	}

	// Type assertion — recover concrete type from interface
	var s Shape = Circle{Radius: 3}
	if c, ok := s.(Circle); ok {
		fmt.Printf("It's a circle with radius %.0f\n", c.Radius)
	}

	// Type switch — handle multiple types
	describe := func(i interface{}){
		switch v := i.(type){
		case int: fmt.Println("int:", v)
		case string:  fmt.Println("string:", v)
		case bool: fmt.Println("bool:", v)
		default: fmt.Printf("unknown type: %T\n", v)
		}
	}
	describe(42)
	describe("hello")
	describe(true)
}