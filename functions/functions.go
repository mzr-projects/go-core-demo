package functions

func FirstMethod(a string) string {
	return a + " the variable"
}

func Add(a int, b int) int {
	return a + b
}

// AddAndRemainder : In go we can have functions return multiple values at the same time
//
//	all function calls in go are call by value /*
func AddAndRemainder(a int, b int) (int, int) {
	return a + b, a % b
}
