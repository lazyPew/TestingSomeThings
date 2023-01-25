package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

/*
Marshal и Unmarshal (кодирование и декодирование) данных в формате JSON
в стандартной библиотеке Go реализовано в пакете encoding/json.

Наиболее удобным типом для кодирования / декодирования таких данных
является структура и срез структур, именно его мы и рассмотрим в рамках этого курса.
Но при этом нужно отметить, что в некоторых случаях, когда структура данных
нам не известна, мы можем декодировать данные в типы с использованием интерфейсов:
interface{}, map[string]interface{}, []interface{}, []map[string]interface{},
однако в дальнейшем такой способ потребует использования рефлексии
для анализа таких данных, а это выходит за пределы рассматриваемой темы.

В рассматриваемом пакете мы можем найти 3 функции, позволяющие кодировать / декодировать
данные в байтовый срез, чтобы иметь возможность рассматривать конкретные примеры работы,
давайте сначала рассмотрим эти функции. Начнем с функции для кодирования данных Marshal:
*/
type myStruct struct {
	Name   string
	Age    int
	Status bool
	Values []int
}

func jsonTestFunc1() {
	s := myStruct{
		Name:   "John Connor",
		Age:    35,
		Status: true,
		Values: []int{15, 11, 37},
	}

	// Функция Marshal принимает аргумент типа interface{} (в нашем случае это структура)
	// и возвращает байтовый срез с данными, кодированными в формат JSON.
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", data)
	// {"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}
}

/*
Как мы видим, данные закодированы и мы даже можем их прочитать.
Если же мы хотим получить результат, который лучше подходит именно
для чтения человеком (например для использования в качестве
конфигурационного файла или для отображения информации
на экране компьютера, а не для передачи данных другой программе по сети),
мы можем использовать родственную функцию MarshalIndent.

MarshalIndent похож на Marshal, но применяет отступ (indent)
для форматирования вывода. Каждый элемент JSON в выходных данных
начинается с новой строки, начинающейся с префикса (prefix),
за которым следует один или несколько отступов в соответствии с вложенностью:
*/

func jsonTestFunc2() {

	s := myStruct{
		Name:   "John Connor",
		Age:    35,
		Status: true,
		Values: []int{15, 11, 37},
	}

	data, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", data)
	//{
	//	"Name": "John Connor",
	//	"Age": 35,
	//	"Status": true,
	//	"Values": [
	//		15,
	//		11,
	//		37
	//	]
	//}
}

/*
Ну и в завершении этого шага рассмотрим последнюю из трех функций Unmarshal,
она принимает в качестве аргумента байтовый срез и указатель на объект,
в который требуется декодировать данные. Рассмотрим это на уже знакомом примере:
*/

func jsonTestFunc3() {
	data := []byte(`{"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}`)

	var s myStruct

	if err := json.Unmarshal(data, &s); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v", s) // {John Connor 35 true [15 11 37]}
}

/*
Проверка json на правильность
Мы можем проверить является ли срез байтов форматом json:
*/

type user struct {
	Name     string
	Email    string
	Status   bool
	Language []byte
}

func jsonTestFunc4() {
	m := user{Name: "John Connor", Email: "test email"}

	data, _ := json.Marshal(m)

	data = bytes.Trim(data, "{") // испортим json удалив '{'

	// функция json.Valid возвращает bool, true - если json правильный
	if !json.Valid(data) {
		fmt.Println("invalid json!") // вывод: invalid json!
	}

	fmt.Printf("%s", data)
	// вывод:
	// "Name":"John Connor","Email":
	// "test email","Status":false,"Language":null}
}

func main() {
	fmt.Println("jsonTestFunc 1:")
	jsonTestFunc1()
	fmt.Println("\n===============")
	fmt.Println("jsonTestFunc 2:")
	jsonTestFunc2()
	fmt.Println("\n===============")
	fmt.Println("jsonTestFunc 3:")
	jsonTestFunc3()
	fmt.Println("\n===============")
	fmt.Println("jsonTestFunc 4:")
	jsonTestFunc4()
	fmt.Println("\n===============")
}
