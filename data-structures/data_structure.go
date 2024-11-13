package data_structures

import (
	"fmt"
)

type Person struct {
	Name    string
	Address string
}

/*
Arrays : Fixed length list of things
Slice  : An array that can grow or shrink but every element in the slice must be of same TYPE
*/
func workWithArrays() {
	cards := []string{"man", "you"}
	for i, card := range cards {
		fmt.Printf("index : %d\t, value : %s\n", i, card)
	}
}

func workWithSlices() {
	cards := []string{"me", "you", newCard("My man")}
	/*
		The append function doesn't modify the cards slice just create a new one
	*/
	cards = append(cards, "HOLY")

	for i, card := range cards {
		fmt.Printf("index : %d\t, value : %s\n", i, card)
	}
}

func workWIthMaps() {
	persons := make(map[int]string)
	persons[1] = "me"
	persons[2] = "you"
	persons[3] = "me"
	lengthOfMap := len(persons)

	for key, value := range persons {
		fmt.Printf("key is : %d and value is : %v and length of map is : %d\n", key, value, lengthOfMap)
	}

	if value, exists := persons[2]; exists {
		fmt.Printf("Value exists and is :%s\n", value)
	} else {
		fmt.Printf("Person does not exist.\n")
	}

	personJohn := Person{
		Name:    "John Doe",
		Address: "New York",
	}

	personNick := Person{
		Name:    "Nick Doe",
		Address: "Old York",
	}

	personMap := make(map[string]Person)
	personMap["John"] = personJohn
	personMap["Nick"] = personNick

	for key, value := range personMap {
		fmt.Printf("key is : %s and value is : %v\n", key, value)
	}
}

func newCard(value string) string {
	return value
}
