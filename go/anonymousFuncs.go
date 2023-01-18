package main

import "fmt"

/*
В Go функции являются объектами первого класса,
это значит, что в этом языке программирования
функцию можно передать в качестве аргумента
другой функции или же вернуть функцию в качестве значения.

Рассмотрим передачу функции в качестве аргумента
другой функции на примере Map из уже знакомого
нам пакета strings. Эта функция выглядит так:

	func Map(mapping func(rune) rune, s string) string

Функция Map в качестве первого аргумента получает
функцию вида func (rune) rune, производящей какие-то
действия с символом Unicode и возвращающей
в качестве результата символ Unicode. Из описания этой функции
следует, что переданная в качестве аргумента функция
будет применена к каждому символу строки,
переданной в качестве второго аргумента функции Map,
получившаяся строка будет возвращена в качестве результата.

Создадим такую функцию:

	func invert(r rune) rune {
		// Если буква строчная, то она возвращается заглавной
		if unicode.IsLower(r) {
			return unicode.ToUpper(r)
		}
		// Иначе возвращается строчной
		return unicode.ToLower(r)
	}

А теперь используем ее:

	func ExampleFirstClassFunctionArgument() {
		src := "aBcDeFg"
		test := "AbCdEfG"
		// Обратите внимание, что скобки после имени функции используются только при ее вызове
		src = strings.Map(invert, src)
		fmt.Printf("Инвертированная строка: %s. Результат: %v.\n", src, src == test)
		// Output:
		// Инвертированная строка: AbCdEfG. Результат: true.
	}

Аналогично мы можем вернуть функцию в качестве значения:

	func returnFunction() func(rune) rune {
		return invert
	}
*/

func main() {
	fmt.Println(SumMultipliedValues(func(a, b int) int {
		return a * b
	}, 10, 200, 7))

	fmt.Println(
		func(a, b int) int {
			return a * b
		}(10, 200))

	fmt.Println(SumMultipliedValues(GetMultiplyFunc(), 10, 200, 7))

	x := func(fn func(i int) int, i int) func(int) int { return fn }(func(i int) int { return i + 1 }, 5)

	fmt.Printf("%T", x)

	// func() func(int) int
	// Т.е. это функция, которая возвращает функцию,
	// принимающую аргумент типа int
	// и возвращающая значение того же типа
}

func GetMultiplyFunc() func(int, int) int {
	return Multiply
}

func Multiply(init, mult int) int {
	return init * mult
}

func SumMultipliedValues(multFunc func(int, int) int, first, second, mult int) int {
	return multFunc(first, mult) + multFunc(second, mult)
}

/*
Ранее, объявляя функцию, мы давали этой функции имя.
Такое объявление можно сделать только за пределами других функций
(на уровне пакета).
Объявляя функцию на уровне пакета мы имеем возможность написать
для такой функции необходимые тесты и удобно изменять ее реализацию.
В этом нам в т.ч. помогает определенный уровень изоляции области видимости такой функции.

Однако в ряде случаев нам необходимо выполнить
определенную задачу на месте, возможно предоставив
функции доступ к области видимости вызывающей функции,
как быть в этом случае?
Язык Go позволяет нам использовать анонимные функции в любом выражении.
Литерал такой функции записывается как объявление функции,
но без имени после ключевого слова func.

Давайте немного изменим пример с функцией Map из пакета strings:

	func ExampleFunctionWithoutName() {
		src := "aBcDeFg"
		test := "AbCdEfG"
		// Обратите внимание, что скобки после имени функции используются только при ее вызове
		src = strings.Map(func(r rune) rune {
			if unicode.IsLower(r) {
				return unicode.ToUpper(r)
			}
			return unicode.ToLower(r)
		}, src)
		fmt.Printf("Инвертированная строка: %s. Результат: %v.\n", src, src == test)
		// Output:
		// Инвертированная строка: AbCdEfG. Результат: true.
	}

В этом примере мы передали функции Map в качестве аргумента
анонимную функцию, общий результат работы от этого не изменился.
Анонимные функции могут быть объявлены в другой функции,
присвоены переменной или вызваны на месте:

	func ExampleUseNoNameFunction() {
		// Присваиваем переменной значение анонимной функции
		fn := func(a, b int) int { return a + b }

		// Выполняем анонимную функцию на месте
		// Обратите внимание на использование скобок при вызове функции
		func(a, b int) {
			fmt.Println(a + b)
		}(12, 34)

		fmt.Println(fn(17, 15))

		// Output:
		// 46
		// 32
	}
В примере мы присвоили переменной fn функцию вида func(int, int) int,
затем выполнили другую анонимную функцию, а затем выполнили функцию,
присвоенную переменной fn. Обратите внимание
на использование скобок в примерах - вызов функции требует наличия скобок,
в которых указываются передаваемые функции аргументы
(если аргументы не передаются - скобки пустые).
*/

/*
	func externalFunction() func() {
		text := "TEXT"

		return func() {
			fmt.Println(text)
		}
	}

	func ExampleEnvironment() {
		fn := externalFunction()
		fn()

		// Output:

Когда анонимная функция использует переменные,
объявленные за ее рамками, ее называют замыканием.
В приведенном примере мы объявили именованную функцию
externalFunction, в которой объявляется переменная text.
В качестве значения функция externalFunction
возвращает замыкание - анонимную функцию, которая
имеет доступ к переменной text, объявленной за ее пределами.
Затем мы присвоили переменной fn результат выполнения
функции externalFunction. Теперь fn - функция,
мы вызываем ее, тем самым печатаем значение переменной text.

Усложним пример:
	func ExampleClosure() {
		fn := func() func(int) int {
			count := 0
			return func(i int) int {
				count++
				return count * i
			}
		}()

		for i := 1; i <= 5; i++ {
			fmt.Println(fn(i))
		}

		// Output:
		// 1
		// 4
		// 9
		// 16
		// 25
	}
*/
