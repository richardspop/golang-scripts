package main

import (
	"fmt"
)

func isLess (str1 string, str2 string) bool {
	if len(str1) < len(str2) {
		//fmt.Println("is Less - ", str1, str2, true)
		return true
	} else if len(str1) > len(str2) {
		//fmt.Println("is Less - ", str1, str2, false)
		return false
	} else if str1 < str2 {
		//fmt.Println("is Less - ", str1, str2, true)
		return true
	} else {
		//fmt.Println("is Less - ", str1, str2, false)
		return false
	}
	//fmt.Println("is Less - ", str1, str2, false)
	return false
}

func swap (slice []string, pos1 int, pos2 int) {
	//fmt.Println("swap ", pos1, pos2)
	slice[pos2], slice[pos1] = slice[pos1],  slice[pos2]
	//fmt.Println(slice)
}

func part(slice []string, low int, high int) int {
	pivot := slice[high]
	//fmt.Println("part ", pivot)
	i := low -1

	for j := low; j <= high -1; j++ {
		//fmt.Println("part idx - ", i, j)
		if isLess(slice[j], pivot) {
			i++
			swap(slice, i, j)
		}
	}
	if isLess(slice[high], slice[i + 1]) {
		swap(slice, i + 1, high)
	}
	//fmt.Println(slice)
	return (i + 1)
}

func quicksort(slice []string, low int, high int) {
	//fmt.Println("quick sort ",low, high)
	if low < high {
		pi := part(slice, low, high)
		//fmt.Println("quick sort part - ",pi)

		quicksort(slice, low, pi - 1)
		quicksort(slice, pi + 1, high)
	}
}

func bigSort(count int, slice []string) {
	quicksort(slice, 0, count - 1)
}

func main() {
	var noOfEle int
	fmt.Scanln(&noOfEle)
	ipStr := make([]string, noOfEle)
	for i := 0 ; i < noOfEle; i++ {
		fmt.Scanln(&ipStr[i])
	}
	//fmt.Println(ipStr)
	bigSort(noOfEle, ipStr)
	//fmt.Println(ipStr)
	for i := 0; i < noOfEle; i++ {
		fmt.Println(ipStr[i])
	}
}
