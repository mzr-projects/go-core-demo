package structs

import "fmt"

type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

type ContactInfo struct {
	email   string
	zipCode int
}

func (pointerToPerson *Person) create(firstName string, lastName string, contact ContactInfo) Person {
	/*
		1st method to initialize the struct in go but, it's not recommended because we must be careful
		about the order of firstName and lastName
		person := Person{"JIM", "CARRY"}
	*/
	/*
		2nd method we specify with fields must be initialized with the provided values
	*/
	person := Person{firstName: firstName, lastName: lastName, contact: contact}
	return person
}

func (pointerToPerson *Person) printPerson(person Person) {
	/*
		%+v will print out multiple field names and their values
	*/
	fmt.Printf("%+v\n", person)
}

/*
*Person -> This is a type description, it means we are working with a pointer to a Person
*pointerToPerson -> Give me the value this memory address is pointing at
 */
func (pointerToPerson *Person) updatePerson(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
