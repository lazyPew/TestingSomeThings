package main

import "fmt"

/*
В ходе анализа результатов переписи населения информация была сохранена в объекте типа map:

groupCity := map[int][]string{
	10:   []string{...}, // города с населением 10-99 тыс. человек
	100:  []string{...}, // города с населением 100-999 тыс. человек
	1000: []string{...}, // города с населением 1000 тыс. человек и более
}
При подготовке важного отчета о городах с населением 100-999 тыс. человек был подготовлен другой объект типа map:

cityPopulation := map[string]int{
	город: население города в тысячах человек,
}
Однако база данных с информацией о точной численности населения содержала ошибки, поэтому в cityPopulation в т.ч. была сохранена информация о городах, которые входят в другие группы из groupCity.

Ваша программа имеет доступ к обоим указанным отображениям, требуется исправить cityPopulation, чтобы в ней была сохранена информация только о городах из группы groupCity[100].

Функция main() уже объявлена, доступ к отображениям осуществляется по указанным именам. По результатам выполнения ничего больше делать не требуется, проверка будет осуществлена автоматически.

*/

func main() {
	groupCity := map[int][]string{
		10:   []string{"Moscaov", "Klinton", "kirovsk"}, // города с населением 10-99 тыс. человек
		100:  []string{"CPB", "NVZ", "Briansk"},         // города с населением 100-999 тыс. человек
		1000: []string{"KIpr", "Harkov", "Novosibirsk"}, // города с населением 1000 тыс. человек и более
	}
	cityPopulation := map[string]int{
		"Moscaov":     54,
		"Klinton":     65,
		"kirovsk":     789,
		"CPB":         345,
		"NVZ":         548,
		"Briansk":     579,
		"KIpr":        65468,
		"Harkov":      78645,
		"Novosibirsk": 57498,
	}
	for cityKey, _ := range cityPopulation {
		var flag bool
		for _, groupStr := range groupCity[100] {
			if cityKey == groupStr {
				flag = true
			}
		}
		if !flag {
			delete(cityPopulation, cityKey)
		}
	}
	fmt.Println(cityPopulation)
}
