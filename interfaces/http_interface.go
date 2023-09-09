package interfaces

import (
	"fmt"
	"net/http"
	"os"
)

func start() {
	resp, err := http.Get("https://www.google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println(resp.Body)
}
