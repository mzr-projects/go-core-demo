package interfaces

import "fmt"

// Bot /*
/*
Interface Type : Means I can not create value of it
All the type inside this program (here englishBot and spanishBot), we have an interface type if you have a type
that has getGreeting function, and you return string then you are a member of this "bot" type
*/
type Bot interface {
	GetGreeting() string
}

// EnglishBot /*
/*
Concrete type : Means I can create a value of it
*/
type EnglishBot struct {
	greeting string
}

// GetGreeting /**
/**
This is the EnglishBot struct method and because it is a method with no arg and returns string go implicitly
implements the Bot interface for it under the hood
*/
func (EnglishBot) GetGreeting() string {
	return "Hi, there"
}

// SpanishBot /*
/*
Concrete type : Means I can create a value of it
*/
type SpanishBot struct {
	greeting string
}

// GetGreeting /**
/*
This is the EnglishBot struct method and because it is a method with no arg and returns string go implicitly
implements the Bot interface for it under the hood
*/
func (SpanishBot) GetGreeting() string {
	return "Ola!"
}

func languageBots() {
	eb := EnglishBot{}
	sb := SpanishBot{}

	printGreetings(eb)
	printGreetings(sb)
}

func printGreetings(eb Bot) {
	fmt.Println(eb.GetGreeting())
}
