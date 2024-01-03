package main

import "fmt"

func main() {
	theum := 3927
	fmt.Println(theum)

	for n := 2; n < theum; n++ {
		if theum%n == 0 {
			theum /= n
			fmt.Printf("を%dで割ると、%d\n", n, theum)
		}
	}
}
