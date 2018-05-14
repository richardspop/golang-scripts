package main

import "reflect"
import "fmt"

func main() {
	a := []uint64{1,2}
	b := []uint64{1,2}
	c := []uint64{1,2,3}
	if reflect.DeepEqual(a, b) {
		fmt.Println("Equal")
	}
	if !reflect.DeepEqual(a, c) {
		fmt.Println("not Equal")
	}
}
