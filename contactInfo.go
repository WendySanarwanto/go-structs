package main

import "fmt"

type contactInfo struct {
	email string
	zip int
}

func (_contactInfo contactInfo) print() {
	fmt.Println("Email: \t", _contactInfo.email)
	fmt.Println("Zip: \t", _contactInfo.zip)
}