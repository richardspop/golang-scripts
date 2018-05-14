package main

import (
	"fmt"
)

type private struct {
		some string
}

func somethin() private {
	return private{"hello"}
}

func main() {
	var ptr *private
	if true {
		temp := somethin()
		ptr = &temp
	}
	fmt.Println(ptr.some)

}
