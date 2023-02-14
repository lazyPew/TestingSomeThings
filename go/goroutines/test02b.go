package main

import (
	"fmt"
)

/*
Любой язык программирования рождает определенные шаблоны
разработки, и Go не является исключением
Теперь канал для синхронизации создается самой функцией
myFunc(), полезная работа выполняется в отдельной горутине.
А еще вы могли обратить внимание, что myFunc возвращает
<-chan struct{}. Стрелка слева от ключевого слова chan означает,
что возвращаемый канал предназначен только для чтения из него.
Аналогичным образом мы можем вернуть из функции или передать
в нее в качестве аргумента канал, предназначенный только для
записи - chan<- struct{}.
*/
func main() {
	<-myFunc()
}

func myFunc() <-chan struct{} {
	done := make(chan struct{})
	go func() {
		fmt.Println("hello")
		close(done)
	}()
	fmt.Println("hell123o")
	return done
}
