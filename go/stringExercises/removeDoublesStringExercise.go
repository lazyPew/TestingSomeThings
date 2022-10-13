package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	for _, s := range str {
		if strings.Count(str, string(s)) > 1 {
			str = strings.Replace(str, string(s), string(""), -1)
		}
	}
	fmt.Println(str)

	// zaabcbd
}
