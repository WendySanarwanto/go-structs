package main

import "fmt"

// Define person struct
type person struct {
	firstName string
	lastName	string
}

func (_person person) print() {
	fmt.Println(_person.firstName, _person.lastName)
}