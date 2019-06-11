package main

import "fmt"

// Define person struct
type person struct {
	firstName string
	lastName	string
	contact		contactInfo
}

func (_person person) print() {
	fmt.Println("Name: \t", _person.firstName, _person.lastName)
	_person.contact.print()
}