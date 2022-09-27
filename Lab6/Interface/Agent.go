package Interface

import "fmt"

type Agent struct {
	FirstName string
	LastName  string
}

func (a Agent) Speak() {
	fmt.Println(a.FirstName, a.LastName)
}
