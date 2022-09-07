package main

import "fmt"

func main() {
	first := 0
	second := 1
	var count, numb int
	fmt.Scan(&numb)
	for count = 1; (first + second) <= numb; count++ {
		first, second = second, (first + second)
		fmt.Println(first, second)
	}
	if numb > second {
		fmt.Println(-1)
	} else {
		fmt.Println(count)
	}
}
