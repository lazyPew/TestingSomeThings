package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
Данная задача поможет вам разобраться в пакете encoding/csv и path/filepath,
хотя для решения может быть использован также пакет archive/zip
(поскольку файл с заданием предоставляется именно в этом формате).

Один из этих файлов в папке task является файлом с данными в формате CSV,
прочие же файлы структурированных данных не содержат.

Требуется найти и прочитать этот единственный файл со структурированными данными
(это таблица 10х10, разделителем является запятая), а в качестве ответа
необходимо указать число, находящееся на 5 строке и 3 позиции (индексы 4 и 2 соответственно).
*/

func readFile(path string, info os.FileInfo, err error) {
	f, err := os.Open(path)
	defer f.Close()
	r := csv.NewReader(f)

	data, err := r.ReadAll()
	if err != nil {
		// Когда мы читаем данные до конца файла io.EOF не возвращается, а служит сигналом к завершению чтения
		// ...
	}
	fn := func(row []string) string {
		return fmt.Sprintf("[\t%s\t]", strings.Join(row, "\t"))
	}
	for numb, row := range data {
		if len(row) > 1 {
			fmt.Printf("file: %s\tRowNumb: %d\tRowValue: %s\t RowLen: %d\n", info.Name(), numb, fn(row), len(row))
		}
	}
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err // Если по какой-то причине мы получили ошибку, проигнорируем эту итерацию
	}
	if info.IsDir() {
		return nil // Проигнорируем директории
	}

	readFile(path, info, err)

	return nil
}

func main() {
	const root = "./task" // Файлы моей программы находятся в другой директории
	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}
}
