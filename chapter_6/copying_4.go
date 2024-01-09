package main

import "fmt"

func main() {
	inta := 5
	adinta := &inta

	bdinta := &inta
	*bdinta = 9

	fmt.Println("*bdintaを変更 ：")
	fmt.Printf("inta の値は %d\n", inta)
	fmt.Printf("adinta の内容は %d\n", *adinta)

}
