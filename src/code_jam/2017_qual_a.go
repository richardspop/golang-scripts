package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
)

//var file_loc = "C:\\Users\\porichar\\Work\\hackerrank\\src\\code_jam\\2017_qual_A-small-practice.in"
//var output_loc = "C:\\Users\\porichar\\Work\\hackerrank\\src\\code_jam\\2017_qual_A-small-practice.out"

var file_loc = "C:\\Users\\porichar\\Work\\hackerrank\\src\\code_jam\\2017_qual_test.in"
var output_loc = "C:\\Users\\porichar\\Work\\hackerrank\\src\\code_jam\\2017_qual_test.out"

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

func parseLastInt(line string) (int) {
	lastInt, err := strconv.Atoi(string(line[len(line)-1]))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	return lastInt
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

	r := bufio.NewReader(f)
	noOfCases := readInt(r)
	for caseCnt := 1; caseCnt <= noOfCases; caseCnt++ {
		pancake_string := readln(r)
		flipper_len := parseLastInt(pancake_string)
		len_of_pancake := len(pancake_string) - 2
		fmt.Println(flipper_len, len_of_pancake)

		//last_flipped := false
		flip_count := 0
		flip_toggle := 0
		for idx, val := range pancake_string[:len_of_pancake] {
			//fmt.Println(idx, string(val))
			curr_head := string(val)
			if flip_toggle != 0 && flip_toggle <= len_of_pancake {
				if curr_head == "+" {
					curr_head = "-"
				} else {
					curr_head = "+"
				}
				flip_toggle++
			}
			if curr_head == "-" && (len_of_pancake - idx) > flipper_len {
				flip_count++
				flip_toggle = 1
				//last_flipped = true
			} else if curr_head == "-"{
				fmt.Println("IMPOSSIBLE")
				break
			}
			fmt.Println(flip_toggle, flip_count)
		}
		fmt.Println(flip_toggle, flip_count)
		fmt.Println("-----------------")
		//if last_flipped == true {
		//	fmt.Println("IMPOSIBLE")
		//}
	}
}
