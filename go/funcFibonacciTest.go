package main

func Fibonacci(n int) int {
	first := 0
	second := 1
	var count int
	for count = 1; count != n; count++ {
		first, second = second, (first + second)
	}
	return second
}
