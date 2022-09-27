package main

import "fmt"

type person struct {
	first string
	last  string
}

func (p person) speak() {
	fmt.Println(p.first, p.last)
}
