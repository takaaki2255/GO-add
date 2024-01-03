package main

import "fmt"

func main() {

	pattern := [3]string{"グー", "チョキ", "パー"}

	fmt.Println("****じゃんけん勝敗リスト****")

	for me := 0; me < 3; me++ {
		fmt.Printf("わたしが%sの時:\n", pattern[me])

		for you := 0; you < 3; you++ {
			score := (3 + me - you) % 3

			if score == 2 {
				fmt.Printf("あなたが%sならわたしの勝ち\n", pattern[you])
			} else if score == 1 {
				fmt.Printf("あなたが%sならわたしの負け\n", pattern[you])
			} else {
				fmt.Printf("あなたが%sならわたしはあいこ\n", pattern[you])
			}
		}
	}
	fmt.Println()
}
