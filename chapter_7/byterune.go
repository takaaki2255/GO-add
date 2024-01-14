package main

import "fmt"

func main() {
	str := "internationalzation"
	strja := "達磨さんが転んだ"

	bt := []byte(str)
	fmt.Println(bt)

	rn := []rune(strja)
	fmt.Println(rn)

	strback := ""
	for i := 0; i < len(bt); i++ {
		strback += string(bt[i])
	}

	strbackja := ""
	for i := 0; i < len(rn); i++ {
		strbackja += string(rn[i])
	}

	fmt.Println(strback)
	fmt.Println(strbackja)
}
