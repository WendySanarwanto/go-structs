package main

import "fmt"

// Define person struct
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p *person) update(_firstName string, _lastName string) {
	(*p).firstName = _firstName
	(*p).lastName = _lastName
}

func (p person) print() {
	fmt.Println("Name: \t", p.firstName, p.lastName)
	p.contact.print()
}
