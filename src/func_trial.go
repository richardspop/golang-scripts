package main

import "fmt"

type test struct {
	a int
}
type trial struct {
	tests []*test
}

func rettest() *trial {
	arg := []*test{}
	a := test{10}
	arg = append(arg, &a)
	t1 := trial{arg}
	return &t1
}

func main() {
	fmt.Println(rettest().tests[0].a)
}
