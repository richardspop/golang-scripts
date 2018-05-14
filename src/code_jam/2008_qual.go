package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

//var file_loc = "C:\\Users\\porichar\\Work\\hackerrank\\src\\code_jam\\2008_qual_test.in"
//var file_loc = "C:\\Users\\porichar\\Work\\hackerrank\\src\\code_jam\\2008_qual_A-small-practice.in"
var file_loc = "C:\\Users\\porichar\\Work\\hackerrank\\src\\code_jam\\2008_qual_A-large-practice.in"
//var output_loc = "C:\\Users\\porichar\\Work\\hackerrank\\src\\code_jam\\2008_qual_A-small-practice.out"
var output_loc = "C:\\Users\\porichar\\Work\\hackerrank\\src\\code_jam\\2008_qual_A-large-practice.out"

func readInt(r *bufio.Reader) (int) {
	it, prefix, err := r.ReadLine()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	} else if prefix == true {
		fmt.Println("prefix is set - not handled")
		os.Exit(-1)
	}
	i , err := strconv.Atoi(string(it))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	return i
}

func readln(r *bufio.Reader) (string) {
	it, prefix, err := r.ReadLine()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	} else if prefix == true {
		fmt.Println("prefix is set - not handled")
		os.Exit(-1)
	}
	return string(it)
}

func main() {
	f, err := os.Open(file_loc)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	defer f.Close()

	w, err := os.Create(output_loc)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	defer w.Close()

	//fmt.Println("file opened")
	r := bufio.NewReader(f)
	noOfCases := readInt(r)
	//fmt.Println(noOfCases)
	for caseCnt := 1; caseCnt <= noOfCases; caseCnt++ {
		noOfSrchEng := readInt(r)
		//fmt.Println(noOfSrchEng)
		srchEng := make(map[string]int)
		for srchCnt := 0; srchCnt < noOfSrchEng; srchCnt ++ {
			srchEng[readln(r)] = 0
			//fmt.Println(srchCnt)
		}
		noOfSrch := readInt(r)
		totSrchEng := len(srchEng)
		tripCnt := 0
		tripAns := 0
		for strCnt := 0; strCnt < noOfSrch; strCnt++ {
			srchStr := readln(r)
			if val, ok := srchEng[srchStr]; ok {
				if val == tripAns {
					srchEng[srchStr]++
					tripCnt++
					if tripCnt == totSrchEng {
						tripAns++
						srchEng[srchStr]++
						tripCnt = 1
						//fmt.Println("tripped at - " + srchStr)
					}
				}
				//fmt.Println("----------------")
				//fmt.Println(srchStr, val)
				//fmt.Println(tripCnt)
				//fmt.Println(tripAns)
				//fmt.Println(srchEng)
				//fmt.Println("----------------")
			}
		}
		fmt.Print("Case #", caseCnt, ": ", tripAns, "\n")
		fmt.Fprint(w, "Case #" + strconv.Itoa(caseCnt) + ": " + strconv.Itoa(tripAns) + "\n")
	}
}
