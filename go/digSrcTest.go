package main

import "fmt"

func main() {
	var numb string
	fmt.Println("start")
	var result int
	fmt.Scan(&numb)
	for numbLen := len(numb); numbLen > 1; numbLen = len(numb) {
		result = 0
		for i := 0; i < numbLen; i++ {
			result += int(numb[i] - '0')
		}
		numb = fmt.Sprintf("%d", result)
	}
	fmt.Println(numb)

}
