package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

/*
Данная задача в основном ориентирована на изучение типа bufio.Reader,
поскольку этот тип позволяет считывать данные постепенно.

В тестовом файле task.data содержится длинный ряд чисел, разделенных символом ";".
Требуется найти, на какой позиции находится число 0 и указать её в качестве ответа.
Требуется вывести именно позицию числа, а не индекс (то-есть порядковый номер, нумерация с 1).

Например:  12;234;6;0;78 :
Правильный ответ будет порядковый номер числа: 4
*/

func walkFunc(path string, info os.FileInfo, err error) error {
	if info.Name() != "task.data" {
		return nil
	}
	f, err := os.Open(path)
	defer f.Close()

	var i int = 0
	r := bufio.NewReader(f)
	for {
		str, _ := r.ReadString(';')
		i++
		if str == "0;" {
			fmt.Printf("%v:\t%s, %[2]T\n", i, str)
			break
		}
	}
	if err != nil {
		panic(err)
	}

	return nil
}

func main() {
	const root = "." // Файлы моей программы находятся в другой директории

	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}
}
