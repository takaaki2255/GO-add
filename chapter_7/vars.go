// package main

// import "fmt"

// func main() {
// 	var something int
// 	var someone string

// 	fmt.Println("something:", something)

// 	if someone == "" {
// 		fmt.Println("someoneは空の文字列")
// 	} else {
// 		fmt.Println("someone:", someone)
// 	}
// }

// package main

// import "fmt"

// type mynumber struct {
// 	value int
// }

// func main() {
// 	var something int
// 	var someone string
// 	var somenumber mynumber

// 	fmt.Println("something:", something)

// 	if someone == "" {
// 		fmt.Println("someoneは空の文字列")
// 	} else {
// 		fmt.Println("someone:", someone)
// 	}
// 	fmt.Println("somenumber:", somenumber)
// }

package main

import "fmt"

type mynumber struct {
	value int
}

func main() {
	var something int
	var someone string
	var somenumber mynumber
	var somearr [5]int
	var somesl []int
	fmt.Println("something:", something)

	if someone == "" {
		fmt.Println("someoneは空の文字列")
	} else {
		fmt.Println("someone:", someone)
	}
	fmt.Println("somenumber:", somenumber)
	fmt.Println("somearr:", somearr)
	fmt.Println("somesl:", somesl)
	if somesl == nil {
		fmt.Println("空のスライスはnil")
	}
}
