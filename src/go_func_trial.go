package main

import "sync"
import "fmt"
import "time"

func main() {
	var wg sync.WaitGroup

	start := time.Now()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1")
		time.Sleep(1 * time.Second)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2")
		time.Sleep(1 * time.Second)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("3")
		time.Sleep(1 * time.Second)
	}()

	wg.Wait()
	fmt.Println("Done - ", time.Now().Sub(start))
}
