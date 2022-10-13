package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	strR := "Строка"
	strE := "String"
	bsR := []byte("Строка")
	bsE := []byte("String")
	rsR := []rune(`Строка`)
	rsE := []rune(`String`)

	fmt.Println("string rus")
	for _, s := range strR {
		fmt.Printf("%v ", s)
	}
	fmt.Println("\nstring eng")
	for _, s := range strE {
		fmt.Printf("%v ", s)
	}

	fmt.Println("\nbyte rus")
	for b := range bsR {
		fmt.Printf("%v ", bsR[b])
	}
	fmt.Println("\nbyte eng")
	for b := range bsE {
		fmt.Printf("%v ", bsE[b])
	}

	fmt.Println("\nrune rus")
	for r := range rsR {
		fmt.Printf("%v ", rsR[r])
	}
	fmt.Println("\nrune eng")
	for r := range rsE {
		fmt.Printf("%v ", rsE[r])
	}
	fmt.Println()

	// len выводит количество байт,
	// латиница и спец символы занимают 1 байт
	// кириллица - 2 байта
	fmt.Println(len(strR), len(strE))
	fmt.Println(utf8.RuneCountInString(strR), utf8.RuneCountInString(strE))

	fmt.Println(string(unicode.ToLower('A')))
	fmt.Println(string(unicode.ToUpper('a')))
	fmt.Println(unicode.IsDigit('a'))
	fmt.Println(unicode.IsDigit('2'))
	fmt.Println(unicode.IsUpper('a'))
	fmt.Println(unicode.IsLower('a'))
	fmt.Println(unicode.Is(unicode.Latin, 'a'))
	fmt.Println(unicode.Is(unicode.Latin, '1'))
	fmt.Println(unicode.Is(unicode.Latin, 'б'))
	fmt.Println()

	fmt.Println(unicode.IsSpace('a'))
	fmt.Println(unicode.IsSpace('\t'))
	fmt.Println(unicode.IsSpace('\v'))
	fmt.Println(unicode.IsSpace('\f'))
	fmt.Println(unicode.IsSpace('\r'))
	fmt.Println(unicode.IsSpace('\n'))
}
