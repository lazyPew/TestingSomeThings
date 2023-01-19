package main

import "fmt"

/*
Приведение типа позволяет нам получить внутреннее значение интерфейса,
в полной мере это реализуется следующим образом:
*/
func testFunc() {
	var i interface{} = 12

	if v, ok := i.(int); ok {
		fmt.Println(v + 12) // Суммирование не произойдет, если ok == false
	}
}

/*
Здесь t значение, приведенное к типу T, а i - исходное значение интерфейсного типа.
ok - логическое значение, показывающее успешность приведения типа:
true - все удачно, false - использование t приведет к панике.
Значение ok может быть опущено, но будьте аккуратны с этим.

Но что делать, если мы не знаем типа объекта, лежащего внутри интерфейса?
Тогда мы можем воспользоваться конструкцией,
называемой переключателем типов, выглядит она следующим образом:
*/
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Умножим на 2:", v*2)
	case string:
		fmt.Println(v + " golang")
	default:
		fmt.Printf("Я не знаю такого типа %T!\n", v)
	}
}

func main() {
	testFunc()

	do(21)
	do("hello")
	do(true)
}
