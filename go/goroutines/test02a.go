package main

import (
	"fmt"
) /*
Каналы предназначены для передачи данных, но могут быть
использованы не только для этой цели. В частности каналы
используются для синхронизации выполнения горутин.

Мы создаем канал done, который будет использован для
синхронизации (тип канала не важен, но пустая структура не
занимает памяти, поэтому использование такого типа крайне
выгодно). myFunc закрывает этот канал после завершения работы.
После запуска myFunc мы ждем, что канал done вернет нам
какое-то значение. В это время myFunc выполняет свою работу и
закрывает канал. Т.к. канал закрыт, а значений в канале нет,
то Go делает вывод, что done уже ничего не вернет и ожидать
значения из него не нужно - функция main продолжает выполнение.
*/

func main() {
	done := make(chan struct{})
	go myFunc(done)
	<-done
}

func myFunc(done chan struct{}) {
	fmt.Println("hello")
	close(done)
}
