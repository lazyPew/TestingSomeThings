package main

import (
	"fmt"
	"strconv"
	"unicode"
)

/*
	Представьте что вы работаете в большой компании
	где используется модульная архитектура.
	Ваш коллега написал модуль с какой-то логикой
	(вы не знаете) и передает в вашу программу какие-то данные.
	Вы же пишете функцию которая считывает две переменных
	типа string ,  а возвращает число типа int64,
	которое нужно получить сложением этих строк.

	Но не было бы так все просто,
	ведь ваш коллега не пишет на Go, и он зол из-за того,
	что гоферам платят больше. Поэтому он решил подшутить
	над вами и подсунул вам свинью.
	Он придумал вставлять мусор в строки перед
	тем как вызывать вашу функцию.

	Поэтому предварительно вам нужно убрать из них мусор
	и конвертировать в числа. Под мусором имеются ввиду
	лишние символы и спец знаки. Разрешается использовать
	только определенные пакеты: fmt, strconv, unicode, strings,  bytes.
*/

func main() {
	fmt.Println(adding("%^80", "hhhhh20&&&&nd"))
}

func adding(a string, b string) int64 {
	var aValue, bValue string
	for _, ar := range a {
		if unicode.IsDigit(ar) {
			aValue += string(ar)
		}
	}
	for _, br := range b {
		if unicode.IsDigit(br) {
			bValue += string(br)
		}
	}
	aInt, err := strconv.Atoi(aValue)
	if err != nil {
		panic(err)
	}
	bInt, err := strconv.Atoi(bValue)
	if err != nil {
		panic(err)
	}
	return int64(aInt) + int64(bInt)
}
