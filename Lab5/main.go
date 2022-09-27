package main

import "fmt"

func main() {

	// struct

	p1 := person{
		first: "Aylin",
		last:  "Çakır",
	}

	fmt.Println(p1.first, p1.last)

	//struct in assignment.This struct is anonymous struct

	p2 := struct {
		first string
		last  string
	}{
		first: "Ahmet",
		last:  "Ay",
	}

	fmt.Println(p2.first, p2.last)

	//Exercise 1
	per1 := exercisePerson{
		first:      "James",
		last:       "Bond",
		favFlavors: []string{"chocolate", "martini", "rum and coke"},
	}

	per2 := exercisePerson{
		first:      "Miss",
		last:       "Moneypenny",
		favFlavors: []string{"strawberry", "vanilla", "capuccino"},
	}

	fmt.Println(per1.first, per1.last)
	for i, v := range per1.favFlavors {
		fmt.Println("\t", i, v)
	}
	fmt.Println(per2.first, per2.last)
	for i, v := range per2.favFlavors {
		fmt.Println("\t", i, v)
	}

	fmt.Println("----------------")

	//Exercise 2

	m := map[string]exercisePerson{
		per1.last: per1,
		per2.last: per2,
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}

	}

	fmt.Println("----------------")

	//Exercise 3

	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "White",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Black",
		},
		luxury: false,
	}

	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.doors)
	fmt.Println(s.doors)

	fmt.Println("----------------")

	//Exercise 4 -> Anonymous Struct

	s1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q":          777,
			"M":          888,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}

	fmt.Println(s1.first)
	fmt.Println(s1.friends)
	fmt.Println(s1.favDrinks)

	for k, v := range s1.friends {
		fmt.Println(k, v)
	}

	for i, val := range s1.favDrinks {
		fmt.Println(i, val)
	}

}
