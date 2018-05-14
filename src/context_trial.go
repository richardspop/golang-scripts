package main

import (
	"context"
	"fmt"
	"time"
)

func for_check(cnt int, ctxVal context.Context) {
	i := 0
	ctxCancel1, cancelFunc1 := context.WithCancel(ctxVal)
	for {
		if i == cnt {
			cancelFunc1()
		} else {
			i++
			fmt.Println(i)
		}
		select {
		case <-ctxCancel1.Done():
			fmt.Println("ctx done. return")
			return
		case <- time.After(1 * time.Second):
			fmt.Println("wait for 1")
		}
	}
}

func startCtx(val int, ctx context.Context) {
	fmt.Println("Ctx started for ", val)
	select {
	case <- ctx.Done():
	}
	fmt.Println("Ctx done for ", val)
}

func main() {
	ctxVal := context.WithValue(context.Background(), 1, "one")
	//ctxCancel, cancelFunc := context.WithCancel(ctxVal)
	//go func() {
	//	fmt.Println("Inside goroutine")
	//	select {
	//	case <-ctxCancel.Done():
	//		fmt.Println("context cancelled")
	//	}
	//	fmt.Println("out of ctx")
	//}()
	//fmt.Println("sleeping for 5 sec")
	//time.Sleep(5 * time.Second)
	//cancelFunc()
	//fmt.Println("canceled")
	//time.Sleep(1 * time.Second)
	//for_check(5, ctxVal)
	ctxCancel0, cancelFunc0 := context.WithCancel(ctxVal)
	ctxCancel1, _ := context.WithCancel(ctxCancel0)
	go startCtx(1, ctxCancel1)
	ctxCancel2, cancelFunc2 := context.WithCancel(ctxCancel0)
	go startCtx(2, ctxCancel2)
	ctxCancel3, _ := context.WithCancel(ctxCancel0)
	go startCtx(3, ctxCancel3)
	cancelFunc2()
	time.Sleep(1 * time.Second)
	cancelFunc0()
	time.Sleep(2 * time.Second)
	go startCtx(3, ctxCancel3)
	time.Sleep(2 * time.Second)
}
