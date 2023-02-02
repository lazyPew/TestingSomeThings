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
даты и времени в следующем формате:
	1986-04-16T05:20:00+06:00
Ваша задача конвертировать эту строку в Time, а затем
вывести в формате UnixDate:
	Wed Apr 16 05:20:00 +0600 1986
Для более эффективной работы рекомендуется ознакомиться с
документацией о содержащихся в модуле time константах.

	Sample Input:
1986-04-16T05:20:00+06:00

	Sample Output:
Wed Apr 16 05:20:00 +0600 1986

*/

func main() {
	// var str string = "1986-04-16T05:20:00+06:00"
	// r := bufio.NewReader(os.Stdin)
	// abs_fname, err := ioutil.ReadAll(os.Stdin)
	// str, _ := r.ReadString('\n')
	// fmt.Println(str)
	buf, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	buf = strings.TrimSuffix(buf, "\n")
	firstTime, err := time.Parse(time.RFC3339, buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(firstTime.Format(time.UnixDate))
}
