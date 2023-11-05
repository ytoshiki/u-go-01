package main 

import "fmt"

type person struct {
	firstName string 
	lastName string
	contact contactInfo
}

type contactInfo struct {
	email string 
	zipCode int
}

func main() {
	var p1 person
	p1.contact = contactInfo{
		email: "test@gmail.com",
		zipCode: 2332222,
	}

	fmt.Printf("%+v", p1)
}