package main

import (
	"fmt"
	"sort"
)



func main() {
	ip := make([]int, 5)
	for i := 0 ; i < 5; i++ {
		fmt.Scan(&ip[i])
	}
	sort.Ints(ip)
	var min, max int64
	min, max = int64(ip[0]), int64(ip[4])
	for i := 1; i < 4; i++ {
		min += int64(ip[i])
		max += int64(ip[i])
	}
	fmt.Println(min, max)
}
