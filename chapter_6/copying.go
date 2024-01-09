package main

import "fmt"

func main() {
	inta := 5
	intb := inta
	inta = 7
	fmt.Printf("inta は %d\n", inta)
	fmt.Printf("intb は %d\n", intb)

	intb = 9
	fmt.Printf("inta は %d\n", inta)
	fmt.Printf("intb は %d\n", intb)
}
