package main

import "fmt"

func dfs(node int, nodes map[int][]int, fn func (int)) {
	dfs_recur(node, nodes, map[int]bool{}, fn)
}

func dfs_recur(node int, nodes map[int][]int, v map[int]bool, fn func (int)) {
	v[node] = true
	fn(node)
	for _, n := range nodes[node] {
		if _, ok := v[n]; !ok {
			dfs_recur(n, nodes, v, fn)
		}
	}
}
func createNode(nodes map[int][]int, city1, city2 int) {
	if list, ok := nodes[city1]; ok {
		list = append(list, city2)
		nodes[city1] = list
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

func searchSlice(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
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
		for i := 0; i < mRoads; i++ {
			var city1, city2 int
			fmt.Scanln(&city1, &city2)
			createNode(nodes, city1, city2)
		}
		//fmt.Println(nodes)
		multiList = append(multiList, nodeList{nodes: nodes, nCities: nCities, mRoads: mRoads, cLib: cLib, cRoad: cRoad})
		//fmt.Println(multiList)
	}
	for _, currList := range multiList{
		nodes := currList.nodes
		nCities := currList.nCities
		cLib := currList.cLib
		cRoad := currList.cRoad
		var cAnsTemp int64
		cAnsTemp = 0
		if cLib <= cRoad {
			cAnsTemp = int64(nCities * cLib)
		} else {
			//fmt.Println(nodes)
			visited := []int{}
			for len(visited) != len(nodes) {
				for city, _ := range nodes {
					if !searchSlice(visited, city) {
						recentVisits := []int{}
						dfs(city, nodes, func(node int) {
							visited = append(visited, node)
							recentVisits = append(recentVisits, node)
						})
						cAnsTemp += int64(cLib + ((len(recentVisits) - 1) * cRoad))
						//fmt.Println(visited, recentVisits)
					}
				}
			}
			cAnsTemp += int64((nCities - len(nodes)) * cLib)
		}
		//fmt.Println(cAnsTemp)
		cAns = append(cAns, cAnsTemp)
	}
	for _, v := range cAns {
		fmt.Println(v)
	}
}
