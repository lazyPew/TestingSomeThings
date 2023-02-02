package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

/*
На стандартный ввод подается строковое представление двух
дат, разделенных запятой (формат данных смотрите в примере).
Необходимо преобразовать полученные данные в тип Time,
а затем вывести продолжительность периода между меньшей
и большей датами.

Sample Input:

	13.03.2018 14:00:15,12.03.2018 14:00:15

Sample Output:

	24h0m0s
*/

func main() {
	// str := "13.03.2018 14:00:15,12.03.2018 14:00:15"
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.TrimSuffix(str, "\n")
	strSplits := strings.Split(str, ",")
	firstTime, err := time.Parse("02.01.2006 15:04:05", strSplits[0])
	if err != nil {
		panic(err)
	}
	secondTime, err := time.Parse("02.01.2006 15:04:05", strSplits[1])
	if err != nil {
		panic(err)
	}
	// fmt.Println(firstTime)
	// fmt.Println(secondTime)
	if firstTime.Before(secondTime) {
		fmt.Println(secondTime.Sub(firstTime))
	} else {
		fmt.Println(firstTime.Sub(secondTime))
	}

}
