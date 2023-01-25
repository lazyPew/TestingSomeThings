package main

import (
	"encoding/json"
	"fmt"
)

func jsonTestFunc5() {
	data1 := []byte(`
	{"menu": {
	"id": "file",
	"value": "File",
	"popup": {
	"menuitem": [
		{"value": "New", "onclick": "CreateNewDoc()"},
		{"value": "!", "onclick": "Doc()"},
		{"value": "Close", "onclick": "CloseDoc()"}
	]
	}		
	}}`)

	data2 := []byte(`
	{"text": {
	"id": "file,
	"value": "kek",
	"popup": {
	"menuitem": [
		{"value": "new2", "power": "CreateNew()"},
		{"value": "new", "onclick": "Close()"}
	]
	}
	}}`)
	fmt.Print(json.Valid(data1))
	fmt.Print(" ")
	fmt.Print(json.Valid(data2))
}

type user struct {
	Name     string
	Email    string
	Status   bool
	Language []byte
}

func jsonTestFunc6() {
	newUser := user{
		Name:     "Alex",
		Email:    "email@email.email",
		Status:   true,
		Language: []byte("ru"),
	}
	data, err := json.Marshal(newUser)
	if err != nil {
		panic(err)
	}
	newUser.Language = []byte("en")
	fmt.Println(string(newUser.Language))
	err = json.Unmarshal(data, &newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(newUser.Language))
}

func main() {
	jsonTestFunc5()
	fmt.Println()
	jsonTestFunc6()

}
