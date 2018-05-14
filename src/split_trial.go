package main

import (
	"strings"
	"fmt"
)

func createSubPath(s []string, to int) string {
	subPath := ""
	for i, sub := range s {
		if i > to {
			return subPath
		} else {
			subPath += "/" + sub
		}
	}
	return subPath
}

func main() {
	path := "/Diag/manifest-agent/whatever"
	s := strings.Split(path, "/")[1:]
	for i, val := range s {
		if val != "" {
			fmt.Println("****")
			fmt.Println(val, path, i)
			fmt.Println(createSubPath(s, i))
		}
	}
	fmt.Println(len(s))
	fmt.Println(s[len(s) -1])

}
