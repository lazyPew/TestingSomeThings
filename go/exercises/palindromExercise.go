package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

/*
	На вход подается строка.
	Нужно определить, является ли она палиндромом.
	Если строка является палиндромом - вывести 'Палиндром'
	иначе - вывести 'Нет'. Все входные строчки в нижнем регистре.

	Палиндром — буквосочетание, слово или текст,
	одинаково читающееся в обоих направлениях (например, "топот", "око", "заказ").
*/

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	lenText := utf8.RuneCountInString(text) - 1
	runeText := []rune(text)

Out:
	for left, s := range runeText {
		right := lenText - 1 - left
		if left >= right {
			fmt.Println("Палиндром")
			break Out
		}

		if s != runeText[right] {
			fmt.Println("Нет")
			break Out
		}
		// топот
	}
}
