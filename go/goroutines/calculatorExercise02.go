package main

import "fmt"

/*
Вам необходимо написать функцию calculator следующего вида:
	func calculator(arguments <-chan int, done <-chan struct{}) <-chan int

В качестве аргумента эта функция получает два канала только
для чтения, возвращает канал только для чтения.
Через канал arguments функция получит ряд чисел, а через
канал done - сигнал о необходимости завершить работу. Когда
сигнал о завершении работы будет получен, функция должна
в выходной (возвращенный) канал отправить сумму полученных чисел.

Функция calculator должна быть неблокирующей, сразу возвращая
управление. Выходной канал должен быть закрыт после выполнения
всех оговоренных условий, если вы этого не сделаете,
то превысите предельное время работы.
*/

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	returnChan := make(chan int)
	x := 0
	go func() {
		// defer close(returnChan)
	L:
		for {
			select {
			case a := <-arguments:
				x += a
				fmt.Println(x)
			case <-done:
				returnChan <- x
				break L
				// close(returnChan)
			}
		}
	}()
	return returnChan
}

func main() {
	ch1 := make(chan int)
	stop := make(chan struct{})

	r := calculator(ch1, stop)
	ch1 <- 3
	ch1 <- 4
	ch1 <- 5
	close(stop)
	fmt.Println(<-r)
}
