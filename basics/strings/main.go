package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	s := "hello, Ayush!"

	//basic ops
	fmt.Println(len(s)) // length in bytes (not runes)
	fmt.Println(strings.ToUpper(s))      // HELLO, AYUSH!
	fmt.Println(strings.Contains(s, "Ay")) // true
	fmt.Println(strings.HasPrefix(s, "he")) //true
	fmt.Println(strings.Replace(s, "hello", "hi", 1))

	// Split and Join
	words := strings.Split("a,b,c", ",")
	fmt.Println(strings.Join(words, " "))

	// TrimSpace — removes leading/trailing whitespace
	fmt.Println(strings.TrimSpace("  hello  "))

	// Efficient string building
	var sb strings.Builder
	for i := range 5{
		sb.WriteString(strconv.Itoa(i)) // no allocations per iteration
		sb.WriteByte(',')
	} 
	fmt.Println(sb.String()) // "0,1,2,3,4,"

    // Convert between string and number
	n := 42
	str := strconv.Itoa(n) // int → string
	fmt.Printf("type: %T value: %s\n", str, str)

	parsed, err := strconv.Atoi("123")
	if err == nil {
		fmt.Println("parsed", parsed) //123
	}

	// Iterating by rune (Unicode-safe)
	for i, r := range "Cafe" {
		fmt.Printf("byte %d: %c\n", i, r)
	}

}	

