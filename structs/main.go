package main 

import "fmt"

type person struct {
	firstName string 
	lastName string
	contactInfo
}

type contactInfo struct {
	email string 
	zipCode int
}

func main() {
	var p1 person
	p1.contactInfo = contactInfo{
		email: "test@gmail.com",
		zipCode: 2332222,
	}

	// &variable gives the memory address of the value the variable is pointing at(not the actual value at this moment)
	// personPointer := &p1
	p1.updateName("Ashley")
	p1.print()
}

// only a pointer to a person type can call the function
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}