package main

import "fmt"

func main() {

	//Exercise 1
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("----------------")

	//Exercise 2
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		fmt.Printf("\t%#U\n", i)

	}

	fmt.Println("----------------")

	//Exercise 3	(While)
	bd := 2001
	for bd <= 2022 {
		fmt.Println(bd)
		bd++
	}

	fmt.Println("----------------")

	//Exercise 4
	bd2 := 2001
	for {
		if bd2 > 2010 {
			break
		}
		fmt.Println(bd2)
		bd2++
	}

	fmt.Println("----------------")

	//Exercise 5
	for i := 10; i <= 100; i++ {
		fmt.Printf("When %v is divided by 4, the remainder is %v\n", i, i%4)
	}

	fmt.Println("----------------")

	//Exercise 6
	x := "James Bond"
	if x == "James Bond" {
		fmt.Println(x)
	}

	fmt.Println("----------------")

	//Exercise 7
	if x == "Moneypenny" {
		fmt.Println(x)
	} else if x == "James Bond" {
		fmt.Println("BONDBONDBONDBONDBONDBOND", x)
	} else {
		fmt.Println("Neither")
	}

	fmt.Println("----------------")

	//Exercise 8
	switch { //If you don't put any variable the default is true
	case false:
		fmt.Println("Should not print")
	case true:
		fmt.Println("Should print")
	}

	fmt.Println("----------------")

	//Exercise 9
	favSport := "Surfing"
	switch favSport {
	case "Skiing":
		fmt.Println("Go to the mountains")
	case "Swimming":
		fmt.Println("Go to the pool!")
	case "Surfing":
		fmt.Println("Go to the Hawaii!")
	case "wingsuit flying":
		fmt.Println("What would you like me to say at your funeral ?")
	}

	fmt.Println("----------------")

	//Exercise 10
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
