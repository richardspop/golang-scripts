package main

import (
	"os"
	"fmt"
	"strconv"
)

var file_loc = "C:\\Users\\porichar\\Work\\share_virtual\\vmr01\\vmr-bundle\\ip_gen.txt"
var ip_prefix = "192.168"
var ip_sub_1  = 123
var ip_sub_2_start = 2
var ip_sub_2_end = 255
var ip_count = 500

func main() {
	w, err := os.Create(file_loc)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	defer w.Close()

	ip_endpoint := ip_sub_2_start

	for cnt := 0; cnt < ip_count; cnt, ip_endpoint = cnt + 1, ip_endpoint + 1 {
		if ip_sub_2_end == ip_endpoint {
			ip_endpoint = ip_sub_2_start
			ip_sub_1++
		}
		fmt.Fprint(w, ip_prefix + "." + strconv.Itoa(ip_sub_1) + "." + strconv.Itoa(ip_endpoint))
		if (cnt + 1) < ip_count {
			fmt.Fprint(w, ",")
		}
	}
}
