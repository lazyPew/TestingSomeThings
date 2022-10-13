package main

import "fmt"

/*
	Оператор defer позволяет выполнить определенную операцию
	после каких-то действий (даже если сработает panic),
	при этом не важно, где в реальности вызывается эта функция.
*/

func main() {

	/*
		Функция finish вызывается с оператором defer,
		поэтому данная функция в реальности будет вызываться
		в самом конце выполнения программы, несмотря на то,
		что ее вызов определен в начале функции main.
	*/
	defer finish()
	defer fmt.Println("Program has been started")
	fmt.Println("Program is working")

	/*
		Команда defer помещает вызов функции в стек.
		Поэтому они выполняются в очередности -LIFO (Last-In, First-Out)

		defer запоминает значения переменных, переданных в функцию,
		на момент объявления defer, а не на момент его вызова.
	*/

	a := 5
	defer printer(a) // когда вызовется printer - будет передано значение 5, а не 7
	a = 7
}

func finish() {
	fmt.Println("Program has been finished")
}
func printer(a int) {
	fmt.Println(a)
}
