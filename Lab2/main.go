package main

import (
	"fmt"
)

func main() {

	//Exercise 1
	a := 42
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)

	fmt.Println("----------------")

	//Exercise 2
	b1 := (42 == 42)
	b2 := (42 <= 43)
	b3 := (42 >= 43)
	b4 := (42 != 43)
	b5 := (42 < 43)
	b6 := (42 > 43)

	fmt.Println(b1, b2, b3, b4, b5, b6)

	fmt.Println("----------------")

	//Exercise 3
	const (
		c     = 42 //Untyped constant
		b int = 42 //Typed constant
	)

	fmt.Println(c, b)

	fmt.Println("----------------")

	//Exercise 4
	d := 42
	fmt.Printf("%d\t%b\t%#x\n", d, d, d)
	e := d << 1 // Burada sağ kısmına 0 ekliyor. Sebebi soldan sağa kayması.
	fmt.Printf("%d\t%b\t%#x\n", e, e, e)
	f := d >> 2 // Ne kadar shift yapacağına buradan karar veriyoruz. >> işareti sağdan sola kaydırıyor.
	fmt.Printf("%d\t%b\t%#x\n", f, f, f)

	fmt.Println("----------------")

	//Exercise 5
	g := `here is something
	as
	a
	raw string
	literal
	"you see"
	another thing`

	fmt.Println(g)

	fmt.Println("----------------")

	//Exercise 6

	const (
		a1 = 2017 + iota
		a2 = 2017 + iota
		a3 = 2017 + iota
		a4 = 2017 + iota
	)

	fmt.Println(a1, a2, a3, a4)
}
