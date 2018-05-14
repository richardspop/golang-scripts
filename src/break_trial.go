package main

import "fmt"

func check_ifelse() {
	a := 10
	defer fmt.Println("end statement")
	if a == 10 {
		fmt.Println("if start")
		return
		fmt.Println("if end")
	} else {
		fmt.Println("else start")
		fmt.Println("end end")
	}
}

func main() {
	check_ifelse()
}
