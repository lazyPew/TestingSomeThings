package main

import (
	"log"
	"time"
)

/*
Внутри функции main, вам необходимо в отдельной горутине
вызвать функцию work() и дождаться результатов ее выполнения.
Функция work() ничего не принимает и не возвращает.
*/

var d string

func main() {

	// *** тут Ваш код ******************

	done := make(chan struct{}) // - создал

	// в анонимку передавать канал не надо.
	// она итак его будет видеть.

	go func() {
		work()
		close(done)
	}() // - передал
	<-done

	// ****************************

	if d != "nil" {
		log.Fatal("Error")
	}
	log.Print("All right")
}

func work() {
	time.Sleep(3 * time.Second)
	d = "nil"
}

/*	1)
<-func() <-chan int {
    c := make(chan int)
    go func() {
        work()
        close(c)
    }()
    return c
}()
*/

/* 	2)
stop := make (chan struct {})

go func (c chan struct {}) {
    work()
    close(c)
}(stop)
<-stop
*/
