package main

import "fmt"

type Person1 struct {
	Name    string
	Address Address
}

type Address struct {
	Number string
	Street string
	City   string
	State  string
	Zip    string
}

func (p *Person1) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func (p *Person1) Location() {
	fmt.Println("Im at", p.Address.Number, p.Address.Street, p.Address.City, p.Address.State, p.Address.Zip)
}

type Citizen struct {
	Country string
	Person1 // анонимное поле без имени
}

func (c *Citizen) Nationality() {
	fmt.Println(c.Name, "is a citizen of", c.Country)
}

func (c *Citizen) Talk() {
	fmt.Println("Hello, my name is", c.Name, "and Im from", c.Country)
}

func main() {
	p := Person1{Name: "Steve"}
	p.Address = Address{Number: "13", Street: "Main"}
	p.Address.City = "Gotham"
	p.Address.State = "NY"
	p.Address.Zip = "01313"
	p.Talk()
	p.Location()

	c := Citizen{}
	c.Name = "Steve"
	c.Country = "America"
	c.Talk()        // <- Метод доступен
	c.Person.Talk() // <- Также доступен
	c.Nationality()
}
