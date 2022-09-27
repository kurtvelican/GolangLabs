package Interface

import "fmt"

type Person struct {
	First string
	Last  string
}

func (p Person) Speak() {
	fmt.Println(p.First, p.Last)
}
