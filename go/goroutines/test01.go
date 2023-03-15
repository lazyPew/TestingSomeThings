package main

import "fmt"

/*
В некоторых случаях горутины могут быть заблокированы.
Это нужно для того, чтобы горутины могли синхронизироваться
между собой.

Горутина, посылающая данные в канал, блокируется, покуда
другая горутина не прочитает данные из него.

Горутина, получающая данные из канала, может быть заблокирована
до момента получения данных в канал. Аналогично блокировке
при записи.

Программе нет смысла использовать системные ресуры в простое -
это затратно по памяти. Для полного понимания, блокировки
горутин можно представить как сделку между двумя людьми:
пока покупатель не заплатит - продавец не предоставит товар.

	Select-case

Используя конструкцию select-case можно избежать блокирующего
поведения.
*/
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		// quit <- 0
	}()
	fibonacci(c, quit)
}

/*
Kод case c <- x: делает запись в канал,
далее считаем новые x и y .

Цикл прокручивается на второй раз и блокируется
(записать второй раз в канал мы уже не можем и
прочитать из канала quit тоже не можем).
Функция fibonacci была запущена внутри горутины main,
и в этом месте эта горутина блокируется.

Это заставляет планировщик запустить анонимную горутину.
И в ней внутри цикла мы из канала читаем записанное
значение 0. Цикл прокрутится второй раз и попробует
еще раз прочитать из канала. И заблокируется при чтении
так как канал без буфера.

Планировщик переключит контекст на горутину main .
В ней внутри fibonacci в select-case мы опять
попадаем в ветку где пишем в канал.
В этот раз пишем 1. И тд.
*/
