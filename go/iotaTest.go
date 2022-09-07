package main

import (
	"fmt"
)

const (
	a = iota + 1
	_
	b
	c
	d = c + 2
	t
	i
	i1
	i2 = iota + 2
	i3 = iota + 2
)

func main() {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	fmt.Println("t = ", t)
	fmt.Println("i = ", i)
	fmt.Println("i1 = ", i1)
	fmt.Println("i2 = ", i2)
	fmt.Println("i3 = ", i3)
}
