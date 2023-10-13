package structs

import (
	"fmt"
	"testing"
)

func TestPerson(t *testing.T) {
	var p = Person{firstName: "Alex", lastName: "Anderson", contact: ContactInfo{email: "aa@gg.com", zipCode: 12344}}
	p.printPerson(p)

	/*
		&p -> Gives me the memory address of the value(p) this variable is pointing at
		pPointer := &p
	*/
	/*
		But instead of using above code pPointer := &p we can simply use the p itself because if we have a receiver
		with the type of pointer(here : (pointerToPerson *Person)) go let us do the call with the pointer like above
		or just with type person. go can do the type conversion
	*/
	p.updatePerson("MyMan!!!")
	p.printPerson(p)
}

func TestStructGotcha(t *testing.T) {
	mySlice := []string{"hi", "there", "How", "are", "you"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

/*
Here the slice will be changed unlike the struct one !!!!!!
Because under the hood the go will create a struct that has length, capacity,pointer to the head of slice

	length	|	cap	 | pointer to the head    ----> Every time we use the slice go use this structure not the
													array inside it, if we pass it to a function it will copy this
													value instead of the value of array

Reference types in go are :
 1. slices
 2. maps
 3. channels
 4. pointers
 5. functions
*/
func updateSlice(s []string) {
	s[0] = "Bye"
}
