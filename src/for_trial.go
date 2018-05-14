package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Open for 1")
		defer fmt.Println("Close for 1")
		for {
			fmt.Println("Open for 2")
			defer fmt.Println("Close for 2")
			time.Sleep(1 * time.Second)
			//continue
			return
		}
	}
}
