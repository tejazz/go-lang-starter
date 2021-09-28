package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Halpert",
		contactInfo: contactInfo{
			email:   "jim.halpert@dunder-mifflin.com",
			zipCode: 94000,
		},
	}

	jimPointer := &jim // & => sets the memory address of the variable
	jimPointer.updateName("Pam")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// *<actual object> => refers to actual value
// *<type> => talks about dealing with pointers
// address => value = *address
// value => address = &value
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
