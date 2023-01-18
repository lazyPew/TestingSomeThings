package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	В привычных нам редакторах электронных таблиц
	присутствует удобное представление числа с разделителем
	разрядов в виде пробела, кроме того в России целая часть
	от дробной отделяется запятой. Набор таких чисел
	был экспортирован в формат CSV, где в качестве разделителя
	используется символ ";".

	На стандартный ввод вы получаете 2 таких вещественных числа,
	в качестве результата требуется вывести частное
	от деления первого числа на второе с точностью
	до четырех знаков после "запятой".

*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _, _ := reader.ReadLine()
	str := string(s)
	var strVals []string
	// str := "1 149,6088607594936;1 179,0666666666666\n"
	str = strings.Split(str, "\n")[0]
	for _, strVal := range strings.Split(str, ";") {
		strVals = append(strVals, strings.Replace(strings.Replace(strVal, " ", "", 1), ",", ".", 1))
	}
	// fmt.Println(strVals)
	aVal, err := strconv.ParseFloat(strVals[0], 64)
	if err != nil {
		panic(err)
	}
	bVal, err := strconv.ParseFloat(strVals[1], 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%.4f\n", aVal/bVal)

	// var a, b float64
	// c, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// c = strings.Replace(c, " ", "", -1)
	// c = strings.Replace(c, ",", ".", -1)
	// fmt.Sscanf(c, "%f;%f", &a, &b)
	// fmt.Printf("%.4f\n", a / b)
}
