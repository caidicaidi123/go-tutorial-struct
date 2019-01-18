package main

import "fmt"

type contactInfo struct {
	email string
	phone int
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	// var alex person
	// alex.firstName = "Alex"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)
	// fmt.Println()
	jim := person{
		"Jim",
		"Party",
		contactInfo{
			"jim@gmail.com",
			87651234,
			123456,
		},
	}
	
	fmt.Printf("%+v", jim)
	fmt.Println()
}