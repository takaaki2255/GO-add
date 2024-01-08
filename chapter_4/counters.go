package main

import "fmt"

func counter() func() {
	ctr := 0
	fmt.Println("カウンタを”初期化しました")
	return func() {
		ctr++
		fmt.Println(ctr)
	}
}

func main() {
	countfunc := counter()
	countfunc()
	countfunc()
	countfunc()
}
