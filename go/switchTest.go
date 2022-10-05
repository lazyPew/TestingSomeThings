package main

import "fmt"

func Switchtest() {

	v := 42
	switch v {
	case 100:
		fmt.Println(100)
		fallthrough
	case 42:
		fmt.Println(42)
		fallthrough
	case 1:
		fmt.Println(1)
		fallthrough
	default:
		fmt.Println("default")
	}

	var c uint32
	fmt.Scan(&c)
	switch {
	case 1 <= c && c <= 9:
		fmt.Println("from 1 to 9")
	case 100 <= c && c <= 250:
		fmt.Println("from 100 to 250")
	case 1000 <= c && c <= 6000:
		fmt.Println("from 1000 to 6000")
	}

	var b float32 = 9
	switch {
	case 1 <= b && b <= 9:
		fmt.Println("from 1 to 9")
		b--
		fallthrough // next block will be done anyway, even if next condition is false
	case b == 10.2:
		fmt.Println("completed")
		fmt.Println(b)
	case 100 <= b && b <= 250:
		fmt.Println("from 100 to 250")
		fmt.Println(b)
	case 1000 <= b && b <= 6000:
		b += 999
		fmt.Println("from 1000 to 6000")
		fallthrough
	default:
		fmt.Println("default")
	}
}
