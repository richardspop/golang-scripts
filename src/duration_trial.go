package main

import (
	"time"
	"fmt"
)

func main() {
	expiryDuration := 24 * time.Hour
	fmt.Println(expiryDuration)
	expiryDuration = 60 * time.Second
	fmt.Println(expiryDuration)
	expiryDuration = 24 * time.Hour
	fmt.Println(expiryDuration)
	var dur time.Duration
	dur = 1
	fmt.Println(dur)
	if dur == 0 {
		fmt.Println("is zero")
	}
}
