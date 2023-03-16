package main

import "fmt"

/*
Вам необходимо написать функцию calculator следующего вида:

	func calculator(firstChan <-chan int, secondChan <-chan int,
		stopChan <-chan struct{}) <-chan int

Функция получает в качестве аргументов 3 канала, и возвращает канал типа <-chan int
  - в случае, если аргумент будет получен из канала firstChan,
    в выходной (возвращенный) канал вы должны отправить
    квадрат аргумента.
  - в случае, если аргумент будет получен из канала secondChan,
    в выходной (возвращенный) канал вы должны отправить
    результат умножения аргумента на 3.
  - в случае, если аргумент будет получен из канала stopChan,
    нужно просто завершить работу функции.

Функция calculator должна быть неблокирующей, сразу возвращая управление.
Ваша функция получит всего одно значение в один из каналов - получили
значение, обработали его, завершили работу.
После завершения работы необходимо освободить ресурсы, закрыв выходной
канал, если вы этого не сделаете, то превысите предельное время работы.
*/

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	returnChan := make(chan int)
	go func() <-chan int {
		defer close(returnChan)
		select {
		case a := <-firstChan:
			returnChan <- (a * a)
			return returnChan
		case b := <-secondChan:
			returnChan <- (b * 3)
			return returnChan
		case <-stopChan:
			return returnChan
		}
	}()
	return returnChan
}

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	stop := make(chan struct{})

	r := calculator(ch1, ch2, stop)
	// ch1 <- 3
	ch2 <- 4
	// close(stop)
	fmt.Println(<-r)
}
