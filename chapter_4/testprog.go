package main

import "fmt"

func assert(fn func(int, int) int,
	a int, b int, resuft int) bool {
	if fn(a, b) == resuft {
		return true
	}
	return false
}

func main() {
	fmt.Println(assert(
		func(a int, b int) int {
			return a + b
		}, 5, 13, 18,
	))

	fmt.Println(assert(
		func(a int, b int) int {
			return a * b * b
		}, 3, 5, 15,
	))
}
