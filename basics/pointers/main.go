package main

import "fmt"

func doubleValue(n int){
	// only changes the local copy
	n *= 2
}

func doublePointer(n *int){
	*n *= 2
}

type Config struct {
	Debug bool
	MaxConn int
}

// Returns a pointer — caller modifies the original
func DefaultConfig() *Config{
	return &Config{Debug: false, MaxConn: 10}
}

func main(){
	x := 5
	doubleValue(5)
	fmt.Println("after doubleValue", x)

	doublePointer(&x)
	fmt.Println("after doublePointer", x)

	// & = address-of,  * = dereference
	p := &x
	fmt.Println("address:", p)
	fmt.Println("value via *:", *p)

	*p = 99
	fmt.Println("x after *p=99", x) //99

	// Pointer to struct
	cfg := DefaultConfig()
	cfg.Debug = true	// Go auto-dereferences: same as (*cfg).Debug = true
	cfg.MaxConn = 50
	fmt.Printf("config: %v\n", *cfg)

  	// new() — allocates zero value
	n := new(int)	// *int pointing to 0
	*n = 42
	fmt.Println("new int:", *n)
}
