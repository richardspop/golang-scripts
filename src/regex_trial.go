package main

import (
	"regexp"
	"fmt"
	"strings"
)

func main() {
	regex := "/api/diag/([^/]+)$"
	match := "http://127.0.0.1:8443/api/diag/"
	re := regexp.MustCompile(regex)
	if re.MatchString(match) {
		fmt.Println("matching", regex, match, re.FindStringSubmatch(match))
	} else {
		fmt.Println("not matching", regex, match, re.FindStringSubmatch(match))
	}
	substr := "/diag"
	if strings.Contains(match, substr) {
		fmt.Println("substring match")
	} else {
		fmt.Println("sub string mismatch")
	}
}
