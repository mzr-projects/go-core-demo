package maps

import (
	"fmt"
	"testing"
)

func TestPerson(t *testing.T) {
	colors := createMap()
	fmt.Println("=== Original Map")
	printNap(colors)
	deleteFromMap(colors)
	fmt.Println("=== After deletion")
	printNap(colors)
}
