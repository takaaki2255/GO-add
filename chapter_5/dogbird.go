package main

import "fmt"

type dog struct {
	name   string
	group  string
	height int
}

type bird struct {
	name   string
	group  string
	height int
}

func (d dog) descride() string {
	dscr := "わたしは" + d.group
	dscr += "、名は" + d.name
	return dscr
}

func (b bird) descride() string {
	dscr := "わたしは" + b.name
	dscr += "、" + b.group + "の仲間"
	return dscr
}

func main() {
	pome := dog{"ポメ", "ポメラニアン", 25}
	meji := bird{"メジロ", "スズメ", 12}

	fmt.Println(pome.descride())
	fmt.Println(meji.descride())
}
