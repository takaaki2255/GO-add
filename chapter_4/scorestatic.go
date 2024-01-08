package main

import "fmt"

func scoreavg(scores []int) int {
	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	return sum / len(scores)
}

func sbsqr(a int, b int) int {
	return (a - b) * (a - b)
}

func avgvar(scores []int, avg int) int {
	sum := 0
	for i := 0; i < len(scores); i++ {
		if avg == 0 {
			sum += scores[i]
		} else {
			sum += sbsqr(scores[i], avg)
		}
	}
	return sum / len(scores)
}

func main() {
	mathscores := []int{40, 89, 77, 68, 72, 39, 58, 87, 93, 48, 65, 74, 60}
	fmt.Printf("%d名による数学試験の結果:\n", len(mathscores))
	fmt.Printf("平均点 %d点:\n", scoreavg(mathscores))
	fmt.Printf("分散 %d(点・点):\n", avgvar(mathscores, avgvar(mathscores, 0)))
}
