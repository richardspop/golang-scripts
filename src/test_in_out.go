package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var in_count, a int
	fmt.Println("start")
	fmt.Scanln(&in_count, &a)
	//fmt.Scanln(&a)
	fmt.Println(in_count)
	fmt.Println(a)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
