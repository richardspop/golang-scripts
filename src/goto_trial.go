package main

import "fmt"

func goto_test()  {
	a := 10
	if a == 1 {
		fmt.Println("if start")
		goto end
		fmt.Println("if end")
	}

	fmt.Println("else start")
	fmt.Println("end end")
	end:
	fmt.Println("end statement")
}

func main() {
	goto_test()
}
