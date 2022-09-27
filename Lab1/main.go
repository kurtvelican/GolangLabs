package main

import (
	"fmt"
)

func main() {

	//Exercise 1
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println("-------------------")

	//Exercise 2
	var x2 int = 42
	var y2 string = "James Bond"
	var z2 bool = true

	fmt.Println(x2, y2, z2)

	fmt.Println(x2)
	fmt.Println(y2)
	fmt.Println(z2)

	fmt.Println("-------------------")

	//Exercise 3
	s := fmt.Sprintf("%v\t%v\t%v", x2, y2, z2)
	fmt.Println(s)

	fmt.Println("-------------------")

	//Exercise 4
	type myType int
	var x3 myType
	fmt.Println(x3)
	fmt.Printf("%T\n", x3)
	x3 = 42
	fmt.Println(x3)

	fmt.Println("-------------------")

	//Exercise 5
	var y3 int
	y3 = int(x3)
	fmt.Println(y3)
	fmt.Printf("%T\n", y3)

}
