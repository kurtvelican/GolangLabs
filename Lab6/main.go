package main

import (
	"GoLangLabs/Lab6/Functions"
	Callback2 "GoLangLabs/Lab6/Functions/Callback"
	"GoLangLabs/Lab6/Interface"
	"fmt"
	"math"
)

func den(s string) (s2 string) {
	s2 = fmt.Sprint(" Who is it ? It is ", s)
	return s2
}

func sayilar(a ...int) { // ... sonsuz değer girebilmemizi sağlıyor. Değerler slice olarak tutuluyor. (0 or more)
	fmt.Println(a) // Variadic parameter
}

/*func den(s string) (string) { Üstteki ile aynı tek değişen return kısmında tanım yapılmadı.
	s2 := fmt.Sprint(" Who is it ? It is ", s)
	return s2
}*/

func main() {
	//My Codes

	defer sayilar(7, 8, 9, 52, 45, 62) // defer komutu bu satırın mainin en alt satırında çalıştırılmasına yarıyor. Dosyalama işlemlerinde close işlemi için kullanılabilir.
	fmt.Println(den("Veli"))
	fmt.Println("----------------")

	Functions.Sum(5, 7) //Using interface
	Functions.Minus(5, 7)
	Functions.Multiply(5, 7)
	Functions.Divide(5, 7)
	Functions.Write("Selam")

	fmt.Println("----------------")

	h1 := Interface.Person{First: "Veli Can", Last: "Kurt"}
	a1 := Interface.Agent{FirstName: "Ali Can", LastName: "Sarak"}
	h1.Speak()
	a1.Speak()

	Interface.Bar(h1)

	fmt.Println("----------------")

	Functions.NestedFunc2()

	fmt.Println("----------------")

	//Anonymous function

	func() {
		fmt.Println("Anonymous func ran.")
	}() // Bu kısım önemli. Buraya fonksiyona vermek istediğimiz parametreleri yazabiliyoruz.

	func(x int) {
		fmt.Println(x)
	}(42)

	fmt.Println("----------------")

	//func expression

	f := func() {
		fmt.Println("This is my first func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("The year big brother started watching:", x)
	}
	g(1984)

	fmt.Println("----------------")

	//Returning a func

	fo := foo()
	fmt.Printf("%T\n", fo())

	fo2 := fo()
	fmt.Println(fo2)

	//Cleaned Up
	fmt.Println(foo()())

	fmt.Println("----------------")

	//Callback -> Passing a func as an argument

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := Callback2.SumCallback(ii...)
	fmt.Println("All Numbers: ", s2)

	s3 := Callback2.EvenSumCallback(Callback2.SumCallback, ii...) //sumCallback -> callback yapılan yer.
	fmt.Println("Even Numbers: ", s3)

	s4 := Callback2.OddSumCallback(Callback2.SumCallback, ii...)
	fmt.Println("Odd Numbers: ", s4)

	fmt.Println("----------------")

	//Closure -> Verinin nerelerde tanımlanabileceği ve tanımlanamayacağı anlatıldı.

	{
		y := 1
		fmt.Println(y)
	}
	// fmt.Println(y) y üstteki kod bloğunda tanımlandığı için bu kodda tanınmıyor.

	a := incrementor()
	b := incrementor()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

	//Cleaned up version -> fmt.Println(incrementor()())

	fmt.Println("----------------")

	//Recursion
	n := Functions.Factorial(5)
	fmt.Println(n)

	//Exercise 1

	n1 := Foo()
	x, s := Bar()

	fmt.Println(n1, x, s)

	//fmt.Println(Foo())
	//fmt.Println(Bar())

	fmt.Println("----------------")

	//Exercise 2
	ii2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x2 := ex21(ii2...)
	fmt.Println(x2)

	ii3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x3 := ex22(ii3)
	fmt.Println(x3)

	fmt.Println("----------------")

	//Exercise 3

	defer ex3() //Defer kullanıldığında üstteki defer satırı en altta bulunur. Daha sonra yazılan deferler en alttakinin üzerine gelir.
	fmt.Println("----------------")
	//Exercise 4

	p1 := ex4{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p1.speak()
	fmt.Println("----------------")
	//Exercise 5

	c := circle{
		radius: 12.345,
	}

	sq := square{
		length: 15,
	}

	info(c)
	info(sq)
	fmt.Println("----------------")

	//Exercise 6

	func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	fmt.Println("----------------")

	//Exercise 7
	ex7 := func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}

	ex7()
	fmt.Printf("%T\n", ex7)

	fmt.Println("----------------")

	//Exercise 8

	ex81 := ex8()
	fmt.Println(ex81())

	fmt.Println("----------------")

	//Exercise 9
	ex91 := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1]
	}

	ex92 := ex9(ex91, []int{1, 2, 3, 4, 5})
	fmt.Println(ex92)

	fmt.Println("----------------")

	//Exercise 10
	ex101 := ex10()
	fmt.Println(ex101())
	fmt.Println(ex101())
	fmt.Println(ex101())

	ex102 := ex10()
	fmt.Println(ex102())
	fmt.Println(ex102())
	fmt.Println(ex102())
}

func foo() func() int { // func() int kısmı foo() fonksiyonunun return kısmıdır.
	return func() int {
		return 451
	}
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func ex21(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func ex22(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func ex3() {
	defer func() {
		fmt.Println("foo DEFER ran")
	}()
	fmt.Println("foo ran")
}

type ex4 struct {
	first string
	last  string
	age   int
}

func (e ex4) speak() {
	fmt.Println("My name is", e.first, e.last, "and I am", e.age, "years old.")
}

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func ex8() func() int {
	return func() int {
		return 42
	}
}

func ex9(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}

func ex10() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
