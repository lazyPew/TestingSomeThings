package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

/*
На стандартный ввод подается строковое представление
даты и времени определенного события в следующем формате:

	2020-05-15 08:00:00

Если время события до обеда (13-00), то ничего менять
не нужно, достаточно вывести дату на стандартный вывод в
том же формате.
Если же событие должно произойти после обеда, необходимо
перенести его на то же время на следующий день, а затем
вывести на стандартный вывод в том же формате.

Sample Input:

	2020-05-15 08:00:00

Sample Output:

	2020-05-15 08:00:00
*/
func main() {
	// var str string = "2020-05-15 08:00:00"
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.TrimSuffix(str, "\n")
	firstTime, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		panic(err)
	}
	if firstTime.Hour() >= 13 {
		firstTime = firstTime.AddDate(0, 0, 1)
	}
	fmt.Println(firstTime.Format("2006-01-02 15:04:05"))
}
