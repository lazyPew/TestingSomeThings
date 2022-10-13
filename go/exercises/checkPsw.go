package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	var str string
	fmt.Scan(&str)
	if utf8.RuneCountInString(str) < 5 {
		fmt.Println("Wrong password")
		return
	}
	for _, s := range str {
		if !unicode.Is(unicode.Latin, s) && !unicode.IsDigit(s) {
			fmt.Println("Wrong password")
			return
		}
	}
	fmt.Println("Ok")
	// fdsghdfgjsdDD1
}
