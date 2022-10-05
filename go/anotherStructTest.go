package main

import "fmt"

type Human2 interface {
	Talk()
}

type Person2 struct {
	Name    string
	Address Address
}

type Address2 struct {
	Number string
	Street string
	City   string
	State  string
	Zip    string
}

func (p *Person2) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func (p *Person2) Location() {
	fmt.Println("I am at", p.Address.Number, p.Address.Street, p.Address.City, p.Address.State, p.Address.Zip)
}

type Citizen2 struct {
	Country string
	Person2
}

func (c *Citizen2) Nationality() {
	fmt.Println(c.Name, "is a citizen of", c.Country)
}

func SpeakTo(h Human2) {
	h.Talk()
}

func main() {
	p := &Person2{Name: "Dave"}
	c := &Citizen2{Person2: Person2{Name: "Steve"}, Country: "America"}

	SpeakTo(p)
	SpeakTo(c)
}
