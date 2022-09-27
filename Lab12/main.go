package main

import (
	"GoLangLabs/Lab12/dog"
	"GoLangLabs/Lab12/mymath"
	"GoLangLabs/Lab12/quote"
	"GoLangLabs/Lab12/word"
	"fmt"
)

type canine struct {
	name string
	age  int
}

func main() {
	fmt.Println("2 + 3 =", mySum(2, 3))
	fmt.Println("4 + 7 =", mySum(4, 7))
	fmt.Println("5 + 9 =", mySum(5, 9))

	//Exercise 1
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))

	/*Test commands for dog directory:
	go test
	go test -bench . //. tüm dosyaları test eder. Sadece bir dosya test edilmek isteniyorsa dosya adı yazılır.
	go test -cover
	go test -coverprofile c.out
	go tool cover -html c.out
	*/

	//Exercise 2
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}

	/*Test commands for quote and word directories:
	go test
	go test -bench . //. tüm dosyaları test eder. Sadece bir dosya test edilmek isteniyorsa dosya adı yazılır.
	go test -cover
	go test -coverprofile c.out
	go tool cover -html c.out
	*/

	//Exercise 3
	xxi := gen()
	for _, v := range xxi {
		fmt.Println(mymath.CenteredAvg(v))
	}

}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

// gen generates data to pass into CenteredAvg0
func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	e := [][]int{a, b, c, d}
	return e
}
