package main

/*
Допустим, нам необходимо запросить у Вконтакте данные о 100 аккаунтах,
если мы попытаемся просто в цикле запросить все эти данные, то
в первую же секунду получим бан, потому что API Вконтакте допускает
совершать не более 3 запросов в секунду. Используя тикер мы можем
отправлять запрос каждые 350 миллисекунд. Похожего эффекта можно
достичь с помощью функции Sleep, но предоставляя каналы таймеры и
тикеры могут работать в многопоточных приложениях.
*/

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	tick := time.NewTicker(time.Second / 2)
	defer tick.Stop()

	wg := new(sync.WaitGroup)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(i, tick.C, wg)
	}

	wg.Wait()

	/*
	 * worker 1 выполнил работу
	 * worker 5 выполнил работу
	 * worker 3 выполнил работу
	 * worker 4 выполнил работу
	 * worker 2 выполнил работу
	 */
}

func worker(id int, limit <-chan time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	<-limit
	fmt.Printf("worker %d выполнил работу\n", id)
}
