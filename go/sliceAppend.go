package main

import "fmt"

/*
	Функция принимает аргумент a типа interface{},
	но перед указанием типа имеется знак многоточия (…).
	Символ многоточия перед указанием типа указывает,
	что в функцию можно передать произвольное количество параметров указанного типа.

	Знак многоточия слева от указания на тип передаваемого значения
	свидетельствует о возможности передать в функцию
	неопределенное количество аргументов указанного типа,
	внутри функции переданные аргументы будут обработаны
	как срез указанного в объявлении функции типа
*/

func MyPrint(a ...interface{}) {
	for _, elem := range a {
		fmt.Printf("%d ", elem)
	}
}

func ExampleMyPrint() {
	MyPrint(1, 2, 3, 4, 5)

	// Output:
	// 1 2 3 4 5
}
func ExampleExpandSlice1() {
	s := []interface{}{1, 2, 3, 4, 5}

	fmt.Println(s)
	fmt.Println(s...)

	// Output:
	// [1 2 3 4 5]
	// 1 2 3 4 5
}

func ExampleExpandSlice2() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{6, 7, 8, 9, 10}

	// append(slice []Type, elems ...Type) []Type

	// s1 = append(s1, s2) не сработает, т.к. второй и
	// последующие аргументы в нашем случае должны быть int

	// Разворачивание (раскрытие) среза путем указания справа от среза знака многоточия

	s1 = append(s1, s2...)
	fmt.Println(s1)

	// Output:
	// [1 2 3 4 5 6 7 8 9 10]
}

func ExpandSliceTest() {
	ExampleMyPrint()
	ExampleExpandSlice1()
	ExampleExpandSlice2()
}
