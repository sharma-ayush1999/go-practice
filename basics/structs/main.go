package main

import "fmt"

// Define a struct type
type Rectangle struct{
	Width float64
	Height float64
}

// Value receiver — does NOT modify the original
// Used for read-only operations
func (r Rectangle) area() float64{
	return r.Height * r.Width
}

// Value receiver — perimeter
func (r Rectangle) perimeter() float64{
	return 2 * (r.Height + r.Width)
}

// Pointer receiver — CAN modify the original
func (r *Rectangle) Scale(factor float64){
	r.Width *= factor
	r.Height *= factor
}


func main(){
   // Struct literal
	rect := Rectangle{Width: 10, Height: 15}
	fmt.Printf("Area: %.1f\n", rect.area())
	fmt.Printf("Perimeter: %.1f\n", rect.perimeter())
	    
	// Pointer receiver — Go auto-takes address (&rect)
	rect.Scale(2)
	fmt.Printf("New Area: %.1f\n", rect.area())

	
    // Anonymous struct — useful for one-off data shapes
	point := struct{
		X, Y int
	}{X: 3, Y: 7}

	fmt.Println("point:", point)
}