package main

import "fmt"

func main() {
	inta := 5
	adinta := &inta
	fmt.Printf("adinta の値は %p\n", adinta)
	fmt.Printf("inta のアドレスは %p\n", &inta)
	fmt.Printf("adinta の内容は　%d\n", *adinta)
}
