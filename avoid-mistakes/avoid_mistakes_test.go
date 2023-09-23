package avoid_mistakes

import (
	"fmt"
	"testing"
)

func TestIfElse(t *testing.T) {
	ifElse, err := correctWayOfIfElse("")
	fmt.Println(err)
	if err != nil {
		return
	}
	fmt.Println(ifElse)
}
