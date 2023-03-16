package main

import (
	"fmt"
	"sync"
)

const N = 10

func main() {
	m := make(map[int]int)
	n := make(map[int]int)

	wg := &sync.WaitGroup{}
	wg2 := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			fmt.Println("i =", i)
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("len without sending i to goroutine =", len(m))

	wg2.Add(N)
	for j := 0; j < N; j++ {
		go func(i int) {
			defer wg2.Done()
			mu.Lock()
			n[i] = i
			fmt.Println("j =", i)
			mu.Unlock()
		}(j)
	}
	wg2.Wait()
	fmt.Println("len with sending j to goroutine =", len(n))
}
