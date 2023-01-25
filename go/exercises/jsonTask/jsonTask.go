package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Student struct {
	LastName   string `json:"LastName"`
	FirstName  string `json:"FirstName"`
	MiddleName string `json:"MiddleName"`
	Birthday   string `json:"Birthday"`
	Address    string `json:"Address"` //"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
	Phone      string `json:"Phone"`   //"+7(948)709-47-24",
	Rating     []int  `json:"Rating"`  //[1, 2, 3]

}
type Group struct {
	ID       int       `json:"ID"`     //134,
	Number   string    `json:"Number"` //"ИЛМ-1274",
	Year     int       `json:"Year"`   //2,
	Students []Student `json:"Students"`
}

func main() {
	// file, err := os.Open("task.json")
	// if err != nil {
	// 	panic(err)
	// }
	// data, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(data)
	abs_fname, err := ioutil.ReadFile("task.json")
	if err != nil {
		panic(err)
	}
	group := Group{}
	// jsonByte, _ := os.ReadFile(abs_fname)
	json.Unmarshal(abs_fname, &group)
	fmt.Println(group)
}
