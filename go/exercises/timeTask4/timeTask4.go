package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
На стандартный ввод подаются данные о продолжительности
периода (формат приведен в примере). Кроме того, вам дана
дата в формате Unix-Time: 1589570165 в виде константы
типа int64 (наносекунды для целей преобразования равны 0).
Требуется считать данные о продолжительности периода,
преобразовать их в тип Duration, а затем вывести
(в формате UnixDate) дату и время, получившиеся при
добавлении периода к стандартной дате.

Небольшая подсказка: базовую дату необходимо явно
перенести в зону UTC с помощью одноименного метода.

Sample Input:

	12 мин. 13 сек.

Sample Output:

	Fri May 15 19:28:18 UTC 2020
*/

const now = 1589570165

func main() {
	// str := "13.03.2018 14:00:15,12.03.2018 14:00:15"
	// str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// str = strings.TrimSuffix(str, "\n")
	var min, sec int
	fmt.Scanf("%d мин. %d сек.", &min, &sec)
	dur, err := time.ParseDuration(strconv.Itoa(min) + "m" + strconv.Itoa(sec) + "s")
	if err != nil {
		panic(err)
	}
	// fmt.Println(dur)
	fmt.Println(time.Unix(now, 0).Add(time.Duration(dur)).Format("Mon Jan 02 15:04:05 UTC 2006"))
}
