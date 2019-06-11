package main

import "fmt"

func main() {
	adaHuang := person { 
		firstName: "Ada", 
		lastName: "Huang",
		contact: contactInfo {
			email: "ada.huang@umbrella.com",
			zip: 8675309, 
		},
	}
	adaHuang.print()
	adaHuang.update("Lisa", "Wong")
	fmt.Println("\nAfter being updated ...\n")
	adaHuang.print()
}
