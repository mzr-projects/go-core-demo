package maps

import "fmt"

func createMap() map[string]string {
	/*
		Here we define a map with the key of string and the value of string
	*/
	/*
		1st approach :

		colors := map[string]string{
			"red":   "#FF0000",
			"green": "#00FF00",
			"blue":  "#0000FF",
		}
	*/

	/*
		2nd approach : If we do not assign the value to it, go will initialize it with zero value
		var colors map[string]string
	*/

	/*
		3rd approach :
		colors := make(map[string]string)
	*/
	colors := make(map[string]string)

	/*
		4th approach:
		colors["red"] = "#FF0000"
	*/
	colors["red"] = "#FF0000"
	colors["white"] = "#FFFFFF"
	colors["green"] = "#00FF00"

	return colors
}

func deleteFromMap(m map[string]string) {
	delete(m, "white")
}

func printNap(m map[string]string) {
	for key, value := range m {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}
}
