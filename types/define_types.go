package types

import (
	"fmt"
	"runtime"
)

type customType int

var b customType
var c int

var x bool
var f float32

var by byte
var rne rune

var s = "This is my goLang DEMO"

/*
This is how we define constants in Go
*/
const constNumber = 1

const (
	aConst = 1
	bConst = 24.65
	cConst = "C_Const"
	dIota  = iota
	eIota  = iota
)

func ShowCustomType() {

	fmt.Println("=============== Custom Types ========")

	b = 43
	fmt.Printf("b is %d\t, Type of b is : %T\n", b, b)

	/*
		This will convert b to type int and assign it to the c variable
		It's conversion not casting
	*/
	c = int(b)
	fmt.Printf("c is %d\t, Type of c is : %T\n", c, c)
}

func ExploreSomeTypes() {

	fmt.Println("=============== Explore Types in GO ========")

	/*
		Default value of bool is FALSE
	*/
	fmt.Printf("x = %t, type is = %T\n", x, x)
	x = true
	fmt.Printf("x = %t, type is = %T\n", x, x)

	a1 := 7
	b1 := 43
	fmt.Printf("a1 == b1 : %t\n", a1 == b1)
	b1 = 7
	fmt.Printf("a1 == b1 : %t\n", a1 == b1)

	/*
		The default type of float is FLOAT64
	*/
	f1 := 23.546
	fmt.Printf("f type is : %T, a1 type is %T, f1 type is %T\n", f, a1, f1)

	/*
		byte is an alias of uint8
		rune is an alias of int32
	*/
	by = 34
	rne = 4353534
	fmt.Printf("by type is %T, rne type is : %T\n", by, rne)

	fmt.Printf("OS is : %s, the Architecture is : %s\n", runtime.GOOS, runtime.GOARCH)

	rawString := `This is a "raw string"
					literal we must pay attention
					to it.`
	fmt.Printf("This is a raw string : %s\n", rawString)
}

func ExploreStrings() {
	/*
		Strings are immutable in GO
	*/
	fmt.Printf("s is \"%s\" and the type of it : %T\n", s, s)
	fmt.Printf("Convert string (s) into slice of byte []byte(s) : %d and type of it is %T\n", []byte(s), []byte(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}

	for i, v := range s {
		fmt.Println(i, v)
	}
}

func ExploreConstants() {
	fmt.Println("=============== Constants and Iota =========")
	fmt.Printf("const is : %d ,type is %T\n", constNumber, constNumber)
	fmt.Printf("const is : %d ,type is %T\n", aConst, aConst)
	fmt.Printf("const is : %f ,type is %T\n", bConst, bConst)
	fmt.Printf("const is : %s ,type is %T\n", cConst, cConst)
	/*
		iota is an auto increment based on the place of it
	*/
	fmt.Printf("dIota is : %d ,type is %T\n", dIota, dIota)
	fmt.Printf("eIota is : %d ,type is %T\n", eIota, eIota)
}

func NumbExercise() {
	y := (42 == 42)
	m := (42 < 43)
	l := (42 < 32)
	fmt.Printf("y is : %t, m is : %t, l is : %t\n", y, m, l)
}
