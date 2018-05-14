package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	rand.Seed(time.Now().UTC().UnixNano())
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
}
