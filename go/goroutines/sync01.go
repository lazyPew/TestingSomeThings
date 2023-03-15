package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Синхронизация горутин с помощью sync.WaitGroup
Еще одну возможность по синхронизации горутин представляет
использование типа sync.WaitGroup. Этот тип позволяет
определить группу горутин, которые должны выполняться
вместе как одна группа. И можно установить блокировку,
которая приостановит выполнение функции, пока не завершит
выполнение вся группа горутин.
*/

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		wg.Add(1)      // Увеличиваем счетчик горутин в группе
		go work(i, wg) // Вызываем функцию work в отдельной горутине
	}

	wg.Wait() // ожидаем завершения всех горутин в группе
	fmt.Println("Горутины завершили выполнение")
}

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Горутина %d начала выполнение \n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Горутина %d завершила выполнение \n", id)
}

/*
В начале определяем группу в виде переменной wg sync.WaitGroup.
В цикле запускаем функцию work 5 раз в отдельных горутинах,
при этом увеличиваем счетчик горутин в группе с помощью метода
Add. Число, которое передается в метод Add определяет значение
внутреннего счетчика активных элементов.
Чтобы сигнализировать, что элемент группы завершил свое
выполнение, в горутине необходимо вызвать метод Done().
Метод Wait блокирует выполнение функции main до завершения
выполнения всех горутин, входящих в группу - его мы вызываем
после запуска горутин, чтобы дождаться результата их выполнения.
Когда внутренний счетчик активных элементов в группе wg станет
равен 0, функция main будет разблокирована и продолжит свое выполнение.
*/
