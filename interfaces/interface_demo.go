package interfaces

import "fmt"

/*
Interface Type : Means I can not create value of it
All the type inside this program (here englishBot and spanishBot), we have an interface type if you have a type
that has getGreeting function, and you return string then you are a member of this "bot" type
*/
type bot interface {
	getGreeting() string
}

/*
Concrete type : Means I can create a value of it
*/
type englishBot struct {
	greeting string
}

/*
Concrete type
*/
type spanishBot struct {
	greeting string
}

func languageBots() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreetings(eb)
	printGreetings(sb)
}

func printGreetings(eb bot) {
	fmt.Println(eb.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi, there"
}

func (spanishBot) getGreeting() string {
	return "Ola!"
}
