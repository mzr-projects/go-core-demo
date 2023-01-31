package main

import (
	"fmt"
	"rsc.io/quote"
)

/*
go fmt : This command will format the code in the current directory
go fmt ./... : This command will format everything all the way down to the last directory it sees
go run main.go : This command will run the code
go build : This command will build the code and create a new exe file in the current directory
go list -m all : This command will print the current module's dependencies
go get : This command will go to some repo and get some code for us
*/

/*
In Go we have to use the variable we defined in the code. Variables have to have usage in
GoLang
*/
func main() {

	/*
		PrintLn has two outputs that we can use them
	*/
	n, err := fmt.Println("Hello GO!!")
	fmt.Printf("n = %d\n", n)
	fmt.Printf("err = %s\n", err)

	/*
		:= is short declaration operator
	*/
	x := 42
	fmt.Printf("x is : %d\n", x)

	/*
		Here we used = because x is already declared
	*/
	x = 54
	fmt.Printf("Now x is : %d\n", x)

	/*
		Calling methods
	*/
	HelloQuote()
	testLoop()
}

func Hello() string {
	return "Hello GO!!"
}

func HelloQuote() string {
	fmt.Println(quote.Hello())
	return quote.Hello()
}

func testLoop() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
