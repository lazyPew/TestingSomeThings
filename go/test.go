package main

import (
	"fmt"
)

func main() {
	var str string
	result := '0'
	fmt.Scan(&str)
	for _, s := range str {
		if s > result {
			result = s
		}
	}
	fmt.Println(string(result))
}
