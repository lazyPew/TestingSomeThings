package main

import (
	"fmt"
)

/*
	На вход дается строка,
	из нее нужно сделать другую строку,
	оставив только нечетные символы (считая с нуля)
*/

func main() {
	var str, result string
	fmt.Scan(&str)
	for i, s := range str {
		if i%2 == 1 {
			result = result + string(s)
		}
	}
	fmt.Println(result)
	// ihgewlqlkot
}
