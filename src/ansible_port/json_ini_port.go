package main

import (
	"fmt"
	"os"
)

var file_loc = "C:\\Users\\porichar\\Work\\hackerrank\\src\\ansible_port\\cisco-vmr.cfg.json"
var output_loc = "C:\\Users\\porichar\\Work\\hackerrank\\src\\ansible_port\\tmp.json"
var data interface{}

func checkAndCreateFile(location string) (f *os.File, err error) {
	if _, err = os.Stat(location); err == nil {
		err = os.Remove(location)
		if err != nil {
			return nil, err
		}
	}
	f, err = os.Create(location)
	return f, err
}


func main() {
	f, err := os.Open(file_loc)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	defer f.Close()

	w, err := checkAndCreateFile(output_loc)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	defer w.Close()



}
