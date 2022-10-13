package main

import (
	"fmt"
	"math"
)

/*
	На вход подается целое число.
	Необходимо возвести в квадрат каждую цифру числа
	и вывести получившееся число.

	Например, у нас есть число 9119.
	Первая цифра - 9. 9 в квадрате - 81.
	Дальше 1. Единица в квадрате - 1.
	В итоге получаем 811181
*/

func main() {
	var str string
	fmt.Scan(&str)
	for _, s := range str {
		fmt.Print(math.Pow(float64(s-'0'), 2))
	}
	fmt.Println()
	// ihgewlqlkot
}
