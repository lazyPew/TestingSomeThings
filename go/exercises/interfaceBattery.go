package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Давайте используем ваши знания структур, методов и интерфейсов на практике и реализуем объект,
удовлетворяющий интерфейсу fmt.Stringer. Назовем его "Батарейка".

Во-первых, вы должны объявить новый тип, удовлетворяющий интерфейсу fmt.Stringer.
Ваш тип должен предусматривать, что на печати он будет выглядеть так: [      XXXX]
где пробелы - "опустошенная" емкость батареи, а X - "заряженная".

Во-вторых, на стандартный ввод вы получаете строку, состоящую ровно из 10 цифр:
0 или 1 (порядок 0/1 случайный). Ваша задача считать эту строку любым возможным способом
и создать на основе этой строки объект объявленного вами на первом этапе типа: надеюсь,
вы понимаете, что строка символизирует емкость батарейки: 0 - это "опустошенная" часть, а 1 - "заряженная".

В-третьих, созданный вами объект должен называться batteryForTest (использование этого имени обязательно).
*/
type Battery struct {
	Charge string
}

func (bat *Battery) String() string {
	var count int = 0
	for _, b := range bat.Charge {
		if b == '1' {
			count++
		}
	}
	str := strings.Repeat(" ", 10-count) + strings.Repeat("X", count)
	return "[" + str + "]"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _, _ := reader.ReadLine()
	str := string(s)

	batteryForTest := &Battery{Charge: str}
	fmt.Println(batteryForTest)
}
