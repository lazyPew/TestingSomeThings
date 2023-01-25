package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
В завершении мы кратко остановимся на двух пакетах, реализующих функционал создания
и обработки путей к файлам path и path/filepath.
Пакет path использует в качестве разделителя символ "/" (слэш), как следствие,
он может использоваться только в операционных системах с соответствующим разделителем -
unix системах (MaсOS, GNU/Linux и другие). Второй же пакет более универсален
и выбирает разделитель в зависимости от операционной системы:
для Windows используется обратный слэш - "\".

Оба пакета содержат достаточное количество примеров их работы,
мы же остановимся на функции Walk из пакета path/filepath:

	func Walk(root string, walkFn WalkFunc) error

	// root - директория, с которой начинается обход
	// walkFn - функция вида func(path string, info os.FileInfo, err error) error

Итак, что же происходит в работе этой функции: Walk рекурсивно обходит все файлы
и директории начиная с директории root и для каждого файла (а директория - тоже файл)
выполняет функцию walkFn. Давайте рассмотрим работу этой функции на примере:

Я заранее создал набор директорий и файлов:
	.
	├── dir1
	│   ├── file1
	│   └── file2
	├── dir2
	│   └── file3
	└── dir3

		├── file4
		├── file5
		└── file6
Давайте соберем информацию о всех файлах:
*/

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err // Если по какой-то причине мы получили ошибку, проигнорируем эту итерацию
	}

	if info.IsDir() {
		return nil // Проигнорируем директории
	}

	fmt.Printf("Name: %s\tSize: %d byte\tPath: %s\n", info.Name(), info.Size(), path)
	return nil
}

func main() {
	const root = "./exercises" // Файлы моей программы находятся в другой директории

	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}

	// Name: file1     Size: 6 byte    Path: test/dir1/file1
	// Name: file2     Size: 6 byte    Path: test/dir1/file2
	// Name: file3     Size: 6 byte    Path: test/dir2/file3
	// Name: file4     Size: 6 byte    Path: test/dir3/file4
	// Name: file5     Size: 6 byte    Path: test/dir3/file5
	// Name: file6     Size: 6 byte    Path: test/dir3/file6
}

/*
Размер файлов одинаков, но это не ошибка. Как вы увидели, с помощью этого инструмента мы можем обработать файлы, отвечающие определенным требованиям: с определенным именем, размером, правами на файл и пр.
*/
