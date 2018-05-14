package main

import (
	"fmt"
	"sort"
)

var currIndex int
var cacheVal = make(map[int][]int)        //key: value, value: list of idx
var cacheDenom = make(map[int][]int)    //key: cacheIdx, value: denominations
var cacheCnt = make(map[int]int)        //key: cacheIdx, value: count

func getCachedCnt(val int, denomList []int) int {
	return -1
}

func getFirstSmallIdx(val int, denomList []int) int {
	for idx, denomVal := range denomList {
		if val >= denomVal {
			return idx
		}
	}
	return -1
}

func getCount(val int, denomList []int) int {
	//fmt.Println(val, denomList)
	ans, idx := 0, 0
	if val == 0 {
		return 1
	}
	if idx = getFirstSmallIdx(val, denomList); idx == -1 {
		return 0
	} else if len(denomList) == 1 {
		if val % denomList[0] == 0 {
			return 1
		} else {
			return 0
		}
	}

	newList := denomList[(idx+1):]
	highValue := denomList[idx]
	maxNoValue := val / highValue
	for i := 0; i <= maxNoValue; i++ {
		//fmt.Println("recur start - ", val, i, maxNoValue, highValue)
		temp := 0
		temp = getCount(val - (i * highValue), newList)
		ans += temp
		cacheVal[val - (i * highValue)] = append(cacheVal[val - (i * highValue)], currIndex)
		cacheDenom[currIndex] = newList
		cacheCnt[currIndex] = temp
		currIndex++
		//fmt.Println("recur end - ", val, i, maxNoValue, highValue)
	}

	return ans
}

func main() {
	var value, noOfDemon int
	currIndex = 0
	fmt.Scanln(&value, &noOfDemon)
	denomList := make([]int, noOfDemon)
	for i := 0; i < noOfDemon; i++ {
		fmt.Scan(&denomList[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(denomList)))
	//fmt.Println(denomList)
	fmt.Println(getCount(value, denomList))
}
