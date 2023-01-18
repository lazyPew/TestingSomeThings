package main

import (
	"fmt"
	"strconv"
)

/*
Внутри функции main (объявлять ее не нужно) вы должны объявить
функцию вида func(uint) uint, которую внутри функции main
в дальнейшем можно будет вызвать по имени fn.

Функция на вход получает целое положительное число (uint),
возвращает число того же типа в котором отсутствуют нечетные
цифры и цифра 0. Если же получившееся число равно 0, то
возвращается число 100. Цифры в возвращаемом числе имеют
тот же порядок, что и в исходном числе.

Уже импортированы - "fmt" и "strconv", другие пакеты импортировать нельзя.

	Sample Input:
	727178
	Sample Output:
	28
*/

func main() {
	fn := func(init uint) uint {
		var res uint
		for _, sVal := range strconv.FormatUint(uint64(init), 10) {
			if (sVal-'0')%2 == 0 && sVal != '0' {
				res = res*10 + uint(sVal-'0')
			}
		}
		if res == 0 {
			return 100
		}
		return res
	}
	fmt.Println(fn(71112980490))
	fmt.Println(fn(0))
}
