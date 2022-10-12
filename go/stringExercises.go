package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

/*
На вход подается строка. Нужно определить, является ли она правильной или нет.
Правильная строка начинается с заглавной буквы и заканчивается точкой.
Если строка правильная - вывести Right иначе - вывести Wrong.

Маленькая подсказка: fmt.Scan() считывает строку до первого пробела,
вы можете считать полностью строку за раз с помощью bufio:
text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
*/

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	lenText := utf8.RuneCountInString(text) - 1
	fmt.Println(lenText)
	fmt.Println((text[lenText]))
Out:
	for i, s := range text {
		switch i {
		case 0:
			if unicode.IsLower(s) {
				fmt.Println("Wrong1")
				break Out
			}
		case lenText - 1:
			if s != '.' {
				fmt.Println("Wrong2")
				fmt.Println((s))
				break Out
			}
		case lenText:
			fmt.Println("Right")
		}
		// if string(s) != "\n" {
		// 	fmt.Println(string(s))
		// }
	}
}
