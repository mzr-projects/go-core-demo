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

	fmt.Printf("OS is : %s, the Architecture is : %s", runtime.GOOS, runtime.GOARCH)
}
