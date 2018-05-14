package main

import (
	"fmt"
)

var a_ans, b_ans int

func triplets (a int, b int) {
	if a > b {
		a_ans++
	} else if a < b {
		b_ans++
	}
}

func main() {
	//fmt.Println("start")
	var a0, a1, a2, b0, b1, b2 int
	fmt.Scanln(&a0, &a1, &a2)
	fmt.Scanln(&b0, &b1, &b2)
	triplets(a0, b0)
	triplets(a1, b1)
	triplets(a2, b2)
	fmt.Println(a_ans, b_ans)
}
