package main

import (
	"errors"
	"fmt"
)

func divide(x, y int) int {
	if y == 0 {
		panic("division by zero!")
	}
	return x / y
}

func main() {
	var input int
	_, err := fmt.Scan(&input) // функция Scan возвращает два параметра, но нам сейчас важно проверить только ошибку
	if err != nil {
		fmt.Println("Проверьте типы входных параметров")
	} else {
		fmt.Println(divide(input, 5)) //Выведем результат, если ошибок нет
	}

	/*
		Стандартная библиотека предоставляет две встроенные функции для создания ошибок:
		errors.New и fmt.Errorf. Обе эти функции позволяют нам
		указывать настраиваемое сообщение об ошибке,
		которое вы можете отображать вашим пользователям.

		errors.New получает один аргумент —
		сообщение об ошибке в виде строки,
		которую вы можете настроить,
		чтобы предупредить ваших пользователей о том, что пошло не так.
	*/

	errNew := errors.New("my error")
	fmt.Println("", errNew)

	/*
		Будет выполняться оператор panic,
		поскольку второй параметр функции divide равен 0.
		И в этом случае все последующие операции,
		которые идут после этого вызова, не будут выполняться
	*/

	fmt.Println(divide(15, 5))
	fmt.Println(divide(4, 0))
}
