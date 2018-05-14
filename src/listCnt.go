package main

import "fmt"

func main() {
	test := []string{"1", "2", "3"}
	maxOR := uint(64)
	for {
		var ans []string
		if uint(len(test)) <= maxOR {
			ans = test
			for i := uint(len(test)); i < maxOR; i++ {
				ans = append(ans, test[len(test)-1])
			}
			test = nil
		} else {
			ans = test[:maxOR]
			test = test[maxOR:]
		}
		if test == nil {
			print(len(ans))
			fmt.Print(ans)
			break
		}
	}
}
