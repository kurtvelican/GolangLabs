package main

import "fmt"

type agent struct {
	firstName string
	lastName  string
}

func (a agent) speak() {
	fmt.Println(a.firstName, a.lastName)
}
