package main

/*
На стандартный ввод подаются данные о студентах
университетской группы в формате JSON:
	{
		"ID":134,
		"Number":"ИЛМ-1274",
		"Year":2,
		"Students":[
			{
				"LastName":"Вещий",
				"FirstName":"Лифон",
				"MiddleName":"Вениаминович",
				"Birthday":"4апреля1970года",
				"Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
				"Phone":"+7(948)709-47-24",
				"Rating":[1,2,3]
			},
			{
				// ...
			}
		]
	}
В сведениях о каждом студенте содержится информация о
полученных им оценках (Rating). Требуется прочитать данные,
и рассчитать среднее количество оценок, полученное
студентами группы. Ответ на задачу требуется записать
на стандартный вывод в формате JSON в следующей форме:
	{
		"Average": 14.1
	}
Как вы понимаете, для декодирования используется функция
Unmarshal, а для кодирования MarshalIndent
(префикс - пустая строка, отступ - 4 пробела).
Если у вас возникли проблемы с чтением / записью данных,
то этот комментарий для вас: в уроках об интерфейсах и
работе с файлами мы рассказывали, что стандартный
ввод / вывод - это файлы, к которым можно обратиться через
os.Stdin и os.Stdout соответственно, они удовлетворяют
интерфейсам io.Reader и io.Writer, из них можно читать,
в них можно писать.

Один из способов чтения, использовать ioutil.ReadAll:
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		// ...
	}

	// data - тип []byte
Задачу можно выполнить и другими способами, в частности
использовать bufio. Буквально в следующем шаге
(через один, на самом деле) будет рассказано о еще одном
способе чтения / записи, можете забежать немного вперед,
а затем вернуться к задаче.
*/
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	LastName   string `json:"-"`
	FirstName  string `json:"-"`
	MiddleName string `json:"-"`
	Birthday   string `json:"-"`
	Address    string `json:"-"`      //"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
	Phone      string `json:"-"`      //"+7(948)709-47-24",
	Rating     []int  `json:"Rating"` //[1, 2, 3]

}
type Group struct {
	ID       int       `json:"-"` //134,
	Number   string    `json:"-"` //"ИЛМ-1274",
	Year     int       `json:"-"` //2,
	Students []Student `json:"Students"`
}

type Average struct {
	Average float32
}

func main() {
	abs_fname, err := ioutil.ReadAll(os.Stdin)
	// abs_fname, err := ioutil.ReadFile("task.json")
	if err != nil {
		panic(err)
	}
	group := Group{}
	json.Unmarshal(abs_fname, &group)
	var studentsCount, ratingCount int
	for _, stud := range group.Students {
		studentsCount++
		ratingCount += len(stud.Rating)
	}
	aver := Average{Average: float32(ratingCount) / float32(studentsCount)}
	data, err := json.MarshalIndent(aver, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", data)
}
