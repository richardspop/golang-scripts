package main

import (
	"fmt"
	"sort"
)

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func mergeSliceReassign(cnt1 int, cnt2 int) {
	l1, l2 := astroMap[cnt1], astroMap[cnt2]
	//j := 0
	initL1 := len(l1)
	for i := 0; i < initL1; i++ {
		if l1[i] != l2[0] {
			if l1[i] > l2[0] {
				l1 = append(l1, l2[0])
				astorList[l2[0]] = cnt1
				//j++
				l2 = remove(l2, 0)
				if len(l2) == 0 {
					break
				}
			}
		} else {
			astorList[l2[0]] = cnt1
			l2 = remove(l2, 0)
			if len(l2) == 0 {
				break
			}
		}
	}
	if len(l2) != 0 {
		for _, val := range l2 {
			l1 = append(l1, val)
			astorList[val] = cnt1
		}
	}
	sort.Ints(l1)
	//fmt.Println(l1, l2)
	astroMap[cnt1] = l1
	delete(astroMap, cnt2)
	//fmt.Println(astorList, astroMap)
}

func addPair(a1 int, a2 int) {
	fmt.Println("add Pair")
	if cnt1, ok := astorList[a1]; ok {
		if cnt2, ok := astorList[a2]; ok {
			if cnt1 != cnt2 {
				//fmt.Println("both astro are present previosuly")
				mergeSliceReassign(cnt1, cnt2)

			} else {
				//fmt.Println("both are from same cnt", a1, cnt1, a2, cnt2)
			}
		} else {
			//fmt.Println("new astro1 ", a2)
			astorList[a2] = cnt1
			astroMap[cnt1] = append(astroMap[cnt1], a2)
		}

	} else if cnt2, ok := astorList[a2]; ok {
		//fmt.Println("new astro2", a1, cnt2)
		astorList[a1] = cnt2
		astroMap[cnt2] = append(astroMap[cnt2], a1)
	} else {
		//fmt.Println("all new astro", a1, a2)
		newPair := []int{a1, a2}
		astroMap[newCntIndex] = newPair
		astorList[a1], astorList[a2] = newCntIndex, newCntIndex
		newCntIndex++
	}
}

var astroMap = make(map[int][]int)	// country : []astronauts
var astorList = make(map[int]int)	// astronauts : country
var newCntIndex int

func main() {
	var noOfAstro, noOfPairs int
	cntIndex := make(map[int]int)
	newCntIndex = int(0)
	fmt.Scanln(&noOfAstro, &noOfPairs)
	for i := int(0); i < noOfPairs; i++ {
		var a1, a2 int
		fmt.Scanln(&a1, &a2)
		addPair(a1, a2)
		//fmt.Println("new list", astroMap, astorList)
	}
	fmt.Println(astorList)
	fmt.Println(astroMap)
	totCnt := len(astroMap) + (noOfAstro - len(astorList))
	//fmt.Println(totCnt, len(astorList))
	for i := 0; i < totCnt; i++ {
		if list, ok := astroMap[i]; ok {
			cntIndex[i] = len(list)
		} else {
			cntIndex[i] = 1
		}
	}
	fmt.Println(cntIndex)
	var ansPerm int64
	for i := 0; i < (totCnt - 1); i++ {
		val1 := int64(cntIndex[i])
		for j := (i + 1); j < (totCnt); j++ {
			val2 := int64(cntIndex[j])
			ansPerm += val1 * val2
			//fmt.Println(ansPerm, val1, val2)
		}
	}
	fmt.Println(ansPerm)

}
