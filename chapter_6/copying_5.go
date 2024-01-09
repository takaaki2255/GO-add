package main

import "fmt"

func main() {
	inta := 5
	adinta := &inta

	bdinta := &inta
	*bdinta = 9

	fmt.Printf("adinta の内容は %p\n", &adinta)
	fmt.Printf("bdinta の内容は %p\n", &bdinta)

}
