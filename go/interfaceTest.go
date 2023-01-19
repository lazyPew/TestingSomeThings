package main

import "fmt"

type calculator interface {
	adderMulter
	subtract(float32, float32) float32
	div(float64, float64) float64
}

type adder interface {
	add(float64, float64) float64
}

type adderMulter interface {
	add(float64, float64) float64
	multi(float64, float64) float64
}

type myCalc struct{}

func (m myCalc) add(a, b float64) float64 {
	return a + b
}

func (m myCalc) subtract(a, b float32) float32 {
	return a - b
}

func (m myCalc) multi(a, b float64) float64 {
	return float64(a * b)
}

func (m myCalc) div(a, b float64) float64 {
	return a / b
}

func main() {
	calc := new(myCalc)
	fmt.Println(calc.add(1, 2))
	fmt.Println(calc.subtract(1, 2))
	fmt.Println(calc.multi(1, 2))
	fmt.Println(calc.div(1, 2))

}
