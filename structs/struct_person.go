package structs

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

type ContactInfo struct {
	email   string
	zipCode int
}

func create(firstName string, lastName string, contact ContactInfo) Person {
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

// createPersonWithPointer is a method that creates a new instance of the Person struct and returns a pointer to it.
// It takes the firstName, lastName and contact parameters as input and initializes the fields of the Person struct.
// The returned pointer points to the newly created Person instance.
func createPersonWithPointer(firstName string, lastName string, contact ContactInfo) *Person {
	return &Person{firstName: firstName, lastName: lastName, contact: contact}
}

/*
Here (pointerToPerson *Person) is the receiver of the create method
*/
func (pointerToPerson *Person) printPerson(person Person) {
	/*
		%+v will print out multiple field names and their values
	*/
	fmt.Printf("%+v\n", person)
}

/*
*Person -> This is a type description, it means we are working with a pointer to a Person
*pointerToPerson -> Gives me the value this memory address is pointing at
 */
func (pointerToPerson *Person) updatePerson(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
