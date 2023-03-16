package main

import (
	"fmt" // пакет используется для проверки выполнения условия задачи, не удаляйте его
	"sync"
	"time"
)

/*
Необходимо написать функцию

	func merge2Channels(fn func(int) int, in1 <-chan int, in2 <- chan int, out chan<- int, n int).

Описание ее работы: n раз сделать следующее
  - прочитать по одному числу из каждого из двух
    каналов in1 и in2, назовем их x1 и x2.
  - вычислить f(x1) + f(x2)
  - записать полученное значение в out

Функция merge2Channels должна быть неблокирующей, сразу возвращая управление.
Функция fn может работать долгое время, ожидая чего-либо или производя вычисления.

Формат ввода:
  - количество итераций передается через аргумент n.
  - целые числа подаются через аргументы-каналы in1 и in2.
  - функция для обработки чисел перед сложением передается через аргумент fn.

Формат вывода: канал для вывода результатов передается через аргумент out.
*/
func main() {
	in1 := make(chan int, 10)
	in2 := make(chan int, 10)
	out := make(chan int, 10)
	n := 10
	start := time.Now()
	for i := 0; i < 10; i++ {
		in1 <- i
		in2 <- 10 + i
	}
	merge2Channels(fn, in1, in2, out, n)
	fmt.Println("0: not blocked if printed first", time.Since(start))
	//
	start = time.Now()
	for i := 0; i < 10; i++ {
		in1 <- i
		in2 <- 10 + i
	}
	merge2Channels1(fn, in1, in2, out, n)
	fmt.Println("1: not blocked if printed first", time.Since(start))
	//
	start = time.Now()
	for i := 0; i < 10; i++ {
		in1 <- i
		in2 <- 10 + i
	}
	merge2Channels2(fn, in1, in2, out, n)
	fmt.Println("2: not blocked if printed first", time.Since(start))
	//
	start = time.Now()
	for i := 0; i < 10; i++ {
		in1 <- i
		in2 <- 10 + i
	}
	merge2Channels3(fn, in1, in2, out, n)
	fmt.Println("3: not blocked if printed first", time.Since(start))
}

func fn(x int) int {
	time.Sleep(time.Millisecond * 100)
	return x
}

var lock sync.Mutex

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
		lock.Lock()

		f1res := make([]*int, n)
		f2res := make([]*int, n)

		input := func(input <-chan int, results []*int) {
			for i := 0; i < n; i++ {
				x := <-input
				go func(i int, x int) {
					res := f(x)
					results[i] = &res
				}(i, x)
			}
		}

		go input(in1, f1res)
		go input(in2, f2res)

		go func() {
			i := 0
			for true {
				if f1res[i] != nil && f2res[i] != nil {
					res := *f1res[i] + *f2res[i]
					out <- res
					if i++; i == n {
						lock.Unlock()
						return
					}
				}
			}
		}()
	}(fn, in1, in2, out, n)
}

// ========================= method 1 =========================

func merge2Channels1(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	c1 := make(chan chan int, n)
	c2 := make(chan chan int, n)
	procIn := func(in <-chan int, c chan chan int) {
		for i := 0; i < n; i++ {
			fc := make(chan int)
			c <- fc
			go func(resChan chan int, x int) {
				resChan <- fn(x)
			}(fc, <-in)
		}
	}
	go procIn(in1, c1)
	go procIn(in2, c2)

	go func() {
		for i := 0; i < n; i++ {
			out <- <-<-c1 + <-<-c2
		}
	}()
}

// ========================= method 2 =========================

type Worker struct {
	sync.Mutex
	sync.WaitGroup
}

func (w *Worker) Do(work func(int) int, x int, res *int) {
	w.Add(1)
	go func() {
		y := work(x)
		w.Lock()
		*res += y
		w.Unlock()
		w.Done()
	}()
}

func merge2Channels2(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		result := make([]int, n)
		var w Worker
		for i := range result {
			w.Do(fn, <-in1, &result[i])
			w.Do(fn, <-in2, &result[i])
		}
		w.Wait()
		for i := range result {
			out <- result[i]
		}
	}()
}

// ========================= method 3 =========================

func merge2Channels3(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		vals := make([]int, n)
		var wg sync.WaitGroup
		wg.Add(2 * n)
		for i := 0; i < n; i++ {
			x1 := <-in1
			x2 := <-in2
			i := i
			go func() {
				vals[i] += fn(x1)
				wg.Done()
			}()
			go func() {
				vals[i] += fn(x2)
				wg.Done()
			}()
		}
		wg.Wait()
		for _, v := range vals {
			out <- v
		}
	}()
}

// ========================= method 4 =========================

func calc_fn(id int, wg *sync.WaitGroup, arr []int, fn func(int) int) {
	defer wg.Done()
	arr[id] = fn(arr[id])
}

func merge2Channels4(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	arr1, arr2 := make([]int, n), make([]int, n)
	wg := new(sync.WaitGroup)
	wg.Add(2 * n)
	// read channels
	go func() {
		for i := 0; i < n; i++ {
			arr1[i] = <-in1
			go calc_fn(i, wg, arr1, fn)
			arr2[i] = <-in2
			go calc_fn(i, wg, arr2, fn)
		}
	}()
	// write to channel
	go func(wg *sync.WaitGroup) {
		wg.Wait()
		for i := 0; i < n; i++ {
			out <- arr1[i] + arr2[i]
		}
	}(wg)
}
