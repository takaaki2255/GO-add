package main

import "fmt"

func main() {
	a := true
	fmt.Printf("aの値は%t\n", a)

	b := false
	fmt.Printf("a==bの値は%t\n", a == b)

	c := 2
	d := 5
	fmt.Printf("c!=dの値は%t\n", c != d)
	fmt.Printf("c<dの値は%t\n", c < d)
	fmt.Printf("(-c)<(-d)の値は%t\n", -1*c < -1*d)

	over1under3 := 1 < c && c < 3
	fmt.Printf("over1under3の値は%t\n", over1under3)
	under5ust5 := d < 5 || d == 5
	fmt.Printf("under5ust5の値は%t\n", under5ust5)
	ctimesdnot10 := (c*d != 10)
	fmt.Printf("ctomesdnot10の値は%t\n", ctimesdnot10)
}
