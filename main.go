/*
We have two kinds of packages in Go lang
 1. Executable (go build -> will build an executable file for us) -> This "go build" command only creates executable for
    the main package other packages get nothing out of it.Package main always must have a function called main
 2. Reusable
*/
package main

import (
	controlflow "com.mt/go-demo/control-flow"
	"com.mt/go-demo/functions"
	"com.mt/go-demo/types"
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
go install : Compiles and installs the packages
go test : Runs any tests associated with the current project
*/

/*
Declare and assign at the same time -> INITIALIZATION
*/
var t = "OMG its outside"
var m = `Jimmy said "BRO it's awesome"`

var xn int = 42
var yn string = "new string"
var zn bool = false

/*
Here we defined a variable with specific type int and assigns the default value of 0 to it
false for booleans
0 for integers
0.0 for floats
" " for strings
nil for pointers,functions, interfaces,slices , channels and maps
*/
var z int

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
		We can also declare variables with "var" keyword
		Teh difference between := and "var" we can use var outside function bodies
	*/
	var y = "This is me"
	fmt.Printf("y Value is : %s , y Type is : %T\n ", y, y)
	fmt.Printf("t Value is : %s , t Type is : %T\n ", t, t)
	fmt.Printf("z Value is : %d , z Type is : %T\n ", z, z)
	fmt.Printf("m Value is : %s , m Type is : %T\n ", m, m)

	s := fmt.Sprintf("%v\t%v\t%v\t", xn, yn, zn)
	fmt.Println(s)

	/*
		Calling methods
	*/
	HelloQuote()
	testLoop()

	fmt.Println("\n=== Types")

	types.ShowCustomType()
	types.ExploreSomeTypes()
	types.ExploreStrings()
	types.ExploreConstants()
	types.NumbExercise()
	types.TypeConversion()

	fmt.Println("=== Control FLow")
	controlflow.TestNestLoop()
	controlflow.TestSomeOtherLoops()
	controlflow.TestForContinue()
	controlflow.TestForRangeLoop()
	controlflow.TestIfElse()
	controlflow.TestSwitch()

	fmt.Println("=== Functions")
	fmt.Println(functions.FirstMethod("SS"))
	fmt.Println(functions.Add(1, 2))
	add, remainder := functions.AddAndRemainder(3, 2)
	fmt.Println("add is :", add, "remainder is:", remainder)
	add, _ = functions.AddAndRemainder(9, 2)
	fmt.Println("add is :", add, "remainder is:", remainder)
	_, remainder = functions.AddAndRemainder(8, 2)
	fmt.Println("add is :", add, "remainder is:", remainder)
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
