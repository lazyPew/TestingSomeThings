package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64 = 10
	var n float64 = 2
	var lim float64 = 123
	if v := math.Pow(x, n); v < lim {
		fmt.Println(v)
		fmt.Println("123")
	}
}
