package control_flow

import (
	"fmt"
	"os"
)

func TestNestLoop() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Outer loop is : %d\t and Inner loop is : %d\n", i, j)
		}
	}
}

func TestSomeOtherLoops() {
	i := 0
	a := 1
	b := 8

	for a < b {
		a *= 2

		fmt.Printf("Number of loops : %d and a is less than 4\n", i)

		if a > 4 {
			break
		}

		i++
	}

	fmt.Printf("DONE !!!!\n")
}

func TestForContinue() {

	x := 1

	for {

		x++
		if x > 10 {
			break
		}

		if x%2 != 0 {
			continue
		}

		fmt.Println(x)
	}

	fmt.Println("DONE we found all even numbers !!")

}

func TestForRangeLoop() {
	//var str string = "Hello, you my man!"
	var str string = "🌍 👋"
	/*
		k is the offset
		v is rune of the string
	*/
	for k, v := range str {
		fmt.Println(k, v, string(v))
	}
}

func TestIfElse() {
	/*
		If we want two statements at the same line we need a ";" for it, like the following
		Now the scope of x is just within the if statement
	*/
	if x := 42; x != 42 {
		fmt.Println("001")
	} else {
		fmt.Println("002 we're in else")
	}
}

func TestSwitch() {
	x := 1
	switch x {
	/*
		We can have multiple matches in case
	*/
	case 1, 3, 4:
		fmt.Println("1 or 3 or 4")
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("Not in our mind ;(")
	}

	/*
		fallthrough will also evaluate the next case if it is true in condition then that
		will be executed as well
	*/
	switch {
	case 1 == 2:
		fmt.Println("1==1")
	case 2 == 2:
		fmt.Println("2==2")
		fallthrough
	case x == 1:
		fmt.Println("x==1")
		fallthrough
	case 3 == 3:
		fmt.Println("3==3")

	}

	word := os.Args[1]
	switch word {
	case "hello":
		fmt.Println("Hi to u")
	case "goodbye":
		fmt.Println("Be safe!")
	case "greetings":
		fmt.Println("HELLLO!!!")
	default:
		fmt.Println("I dont know what to say")

	}

	c := "cracker_joint!"
	switch l := len(word); {
	case word == "hi":
		fmt.Println("Very formal")
	case word == "hello":
		fmt.Println("Hi me")
	case l == 1:
		fmt.Println("I cant understand this one character word ;(")
	case 1 < l && l < 10, word == c:
		fmt.Println("This word is either ", c, " or its 2-9 characters long.")
	default:
		fmt.Println("I dont know the word but it has ", l, " characters.")
	}
}
