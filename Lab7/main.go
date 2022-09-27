package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // &a is the memory address of a. & gives you the address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	var b *int = &a // b := &a is the same as var b *int = &a
	fmt.Println(b)
	fmt.Println(*b) // *b is the value of the memory address of b. * gives you the value stored at an address when you have the address
	fmt.Println(*&a)

	*b = 43
	fmt.Println(a)

	x := 0
	fmt.Println("x before :", &x)
	fmt.Println("x before :", x)
	foo(&x)
	fmt.Println("x after  :", &x)
	fmt.Println("x after  :", x)

	fmt.Println("----------------")

	//Exercise 1
	c := 52
	fmt.Println(c)
	fmt.Println(&c)

	fmt.Println("----------------")

	//Exercise 2

	p1 := person{
		name: "James Bond",
	}
	fmt.Println(p1.name)
	ChangeMe(&p1)
	fmt.Println(p1.name)

}

func foo(y *int) {
	fmt.Println("y before :", y)
	fmt.Println("y before :", *y)
	*y = 43
	fmt.Println("y after  :", y)
	fmt.Println("y after  :", *y)
}
