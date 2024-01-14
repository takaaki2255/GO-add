// package main

// import "fmt"

// func main() {
// 	a := 5 * 2.5
// 	fmt.Println(a)

// 	b := 6
// 	c := b + 0.2
// 	fmt.Println(c)
// }

package main

import "fmt"

func main() {
	a := 5 * 2.5
	fmt.Println(a)

	b := 6
	c := float64(b) + 0.2
	fmt.Println(c)

	fmt.Printf("aのデータ方は%T\n", a)
	fmt.Printf("bのデータ方は%T\n", b)
	fmt.Printf("cのデータ方は%T\n", c)
}
