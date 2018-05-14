package main

import (
	"fmt"
	"time"
)

func main() {
	var a uint32
	var b time.Duration
	a = 2
	fmt.Println(a, b)
	b = time.Duration(a) * time.Minute
	fmt.Println(time.Duration(a), b)
}
