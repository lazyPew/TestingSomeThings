package main

import (
	"fmt"
	"math"
)

func PowPrint(numb, pow, limit float64) {
	var x float64 = numb
	var n float64 = pow
	var lim float64 = limit
	if v := math.Pow(x, n); v < lim {
		fmt.Println(v)
		fmt.Println("123")
	}
}
