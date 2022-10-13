package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/*
	Дана строка, содержащая только английские буквы (большие и маленькие).
	Добавить символ ‘*’ (звездочка) между буквами
	(перед первой буквой и после последней символ ‘*’ добавлять не нужно).

	Входные данные
	Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков.

	Выходные данные
	Вывести строку, которая получится после добавления символов '*'.
*/

func main() {
	var str string
	var count int
	fmt.Scan(&str)
	for i, s := range str {
		if i != 0 && i != utf8.RuneCountInString(str) && string(s) != "*" {
			str = strings.Join([]string{str[:i+count], str[i+count:]}, "*")
			count++
		}
	}
	fmt.Println(str)
}
