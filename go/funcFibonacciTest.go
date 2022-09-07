package main

import "fmt"

func fibonacci(n int) int {
	first := 0
	second := 1
	var count int
	for count = 1; count != n; count++ {
		first, second = second, (first + second)
	}
	return second
}

func main() {
	var numb int
	fmt.Scan(&numb)
	fmt.Println(fibonacci(numb))
}
