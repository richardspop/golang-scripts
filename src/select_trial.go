package main

import (
	"time"
	"fmt"
)
var expir chan bool

func select_exp()  {
	fmt.Println("starting")
	select {
	case <- time.After(time.Until(time.Now().Add(10*time.Second))):
		fmt.Println("timeout")
	case <-  expir:
		fmt.Println("expiry channel")
	}
	fmt.Println("ending")
}

func main() {
	expir = make(chan bool, 1)
	go select_exp()
	time.Sleep(5*time.Second)
	expir <- true
	time.Sleep(20*time.Second)
}
