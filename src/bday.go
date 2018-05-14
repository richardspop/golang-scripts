package main

import "fmt"

func main() {
	var noOfCandles, tallest, height, noOfTallest int
	fmt.Scanln(&noOfCandles)
	tallest = 0
	noOfTallest = 0
	for i := 0; i < noOfCandles; i++ {
		fmt.Scan(&height)
		if height > tallest {
			tallest = height
			noOfTallest = 1
		} else if height == tallest {
			noOfTallest++
		}
	}
	fmt.Println(noOfTallest)
}
