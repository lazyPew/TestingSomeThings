package main

import "fmt"

func fnA(a [3]int) {
	a[1] = 15
}

func fnB(a []int) {
	a[1] = 15
}

func main() {
	baseArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("basearray: %v\n", baseArray)
	baseSlice := baseArray[5:8]
	fmt.Printf("slise based on baseArray: len = %d, capacity = %d\n%v\n", len(baseSlice), cap(baseSlice), baseSlice)
	pointer := fmt.Sprintf("%p", baseSlice)

	fmt.Println("\n============================\n")

	fmt.Println(pointer)
	baseSlice = append(baseSlice, 10)
	fmt.Printf("basearray: %v\n", baseArray)
	fmt.Printf("slise based on baseArray: len = %d, capacity = %d\n%v\n", len(baseSlice), cap(baseSlice), baseSlice)
	fmt.Println(pointer == fmt.Sprintf("%p", baseSlice))

	fmt.Println("\n============================\n")

	fmt.Println(pointer)
	baseSlice = append(baseSlice, 11, 12, 13)
	fmt.Printf("basearray: %v\n", baseArray)
	fmt.Printf("slise based on baseArray: len = %d, capacity = %d \n%v\n", len(baseSlice), cap(baseSlice), baseSlice)
	fmt.Println(pointer == fmt.Sprintf("%p", baseSlice))

	fmt.Println("\n============================\n")

	baseSlice = append(baseSlice[0:2], baseSlice[5:]...)
	fmt.Println(baseSlice)

	fmt.Println("\n============================\n")

	c := []int{4, 5, 6}
	d := make([]int, 2, 2)
	n := copy(d, c)
	fmt.Printf("c = %v\n", c)
	fmt.Printf("d = %v\n", d)
	fmt.Printf("%d elements have been copied\n", n)

	fmt.Println("\n============================\n")

	a := [3]int{1, 2, 3}
	b := []int{1, 2, 3}
	fnA(a)

	/*	fnA(b)
	    fmt.Printf("a = %v\n", a)
	    fmt.Printf("b = %v\n", b)
		fnB(a)*/

	fnB(b)

	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)

}
