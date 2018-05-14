package main

import "fmt"

func isDisJointed (nodes map[int][]int) bool {
	for _, connect := range nodes {
		if len(connect) >= (len(nodes) -1) {
			return false
		}
	}
	return true
}

func createVisited(nodes map[int][]int, city1, city2 int) {
	if list, ok := nodes[city1]; ok {
		list = append(list, city2)
		nodes[city1] = list
		if list2, ok := nodes[city2]; ok {
			list = append(list)
		}
	} else {
		newList := []int{city2}
		nodes[city1] = newList
	}
	if list, ok := nodes[city2]; ok {
		list = append(list, city1)
		nodes[city2] = list
	} else {
		newList := []int{city1}
		nodes[city2] = newList
	}
}

func getHigh(nodes map[int][]int) int {
	lenHigh, cityHigh := 0, 0
	for city, connect := range nodes {
		if len(connect) > lenHigh {
			cityHigh = city
			lenHigh = len(connect)
		}
	}
	fmt.Println("highest - ", cityHigh)
	return cityHigh
}

func removeNodes(nodes map[int][]int, city int) {
	listToDelete := nodes[city]
	delete(nodes, city)
	for _, citytoDel := range listToDelete {
		delete(nodes, citytoDel)
	}
	fmt.Println("after removed ", nodes)
}

type nodeList struct {
	nodes map[int][]int
	nCities int
	mRoads int
	cLib int
	cRoad int
}

func main() {
	var queries int
	cAns := make([]int64, queries)
	fmt.Scanln(&queries)
	var multiList = []nodeList{}
	for ;queries > 0; queries-- {
		var nodes = make(map[int][]int)
		var nCities, mRoads, cLib, cRoad int
		fmt.Scanln(&nCities, &mRoads, &cLib, &cRoad)
		//var visited = make(map[int][]int)
		for i := 0; i < mRoads; i++ {
			var city1, city2 int
			fmt.Scanln(&city1, &city2)
			if cLib > cRoad {
				createVisited(nodes, city1, city2)
			}

		}
		//fmt.Println(nodes)
		multiList = append(multiList, nodeList{nodes: nodes, nCities: nCities, mRoads: mRoads, cLib: cLib, cRoad: cRoad})
		//fmt.Println(multiList)
	}
	for _, currList := range multiList{
		var cAnsTemp int64
		cAnsTemp = 0
		if len(currList.nodes) == 0 {
			//fmt.Println("no nodes")
			cAnsTemp = int64(currList.nCities * currList.cLib)
		} else if !isDisJointed(currList.nodes) {
			cAnsTemp = int64(currList.cLib + ((len(currList.nodes) - 1) * currList.cRoad))
			fmt.Println("it is joined")
			if len(currList.nodes) < currList.nCities {
				cAnsTemp += int64((currList.nCities - len(currList.nodes)) * currList.cLib)

			}
		} else {
			fmt.Println("it is disjointed, ", currList.nodes)
			tempNode := currList.nodes
			city := getHigh(tempNode)
			cAnsTemp += int64(currList.cLib + ((len(tempNode[city])) * currList.cRoad))
			for isDisJointed(tempNode) {
				//fmt.Println("it is disjointed")
				removeNodes(tempNode, city)
				//fmt.Println(tempNode)
				city = getHigh(tempNode)
				cAnsTemp += int64(currList.cLib + ((len(tempNode[city])) * currList.cRoad))
			}
		}
		cAns = append(cAns, cAnsTemp)
	}
	for _, v := range cAns {
		fmt.Println(v)
	}
}
