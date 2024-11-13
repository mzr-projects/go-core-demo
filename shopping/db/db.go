package db

type Item struct {
	/*
		If name starts with lowercase it won't be visible outside this package
	*/
	Name string
}

// LoadItem
// If we use loadItem as the name of the function it won't be visible outside this db package
// /**
func LoadItem(id int) *Item {
	return &Item{Name: "Default"}
}
