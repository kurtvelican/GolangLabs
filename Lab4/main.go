package main

import (
	"fmt"
	"sort"
)

func main() {

	// %#U -> Sayıyı asciiye çevirir.

	dilim := []int{1, 5, 7, 6, 3, 9}
	fmt.Println(dilim)

	// "..." is using for unpack the slice -> dilim2
	dilim2 := []int{52, 25, 78, 96, 43, 51}
	dilim = append(dilim, dilim2...)
	fmt.Println(dilim)

	//sorting slice
	sort.Sort(sort.IntSlice(dilim)) //if you have string slice then use sort.StringSlice()
	fmt.Println(dilim)

	dilim3 := append(make([]int, 10, 100), 57, 8, 96, 35, 45, 48)
	fmt.Println(dilim3)
	dilim3 = append(dilim3, 115, 41, 78, 6, 1, 24, 35)
	fmt.Println(dilim3)

	//map
	m1 := map[int]string{
		1: "Veli",
		2: "Ali",
		3: "İbo",
	}

	fmt.Println(m1[2])

	if v, ok := m1[1]; ok {
		fmt.Println("This is the if print", v)
	}
	//add
	m1[4] = "Ahmet"

	fmt.Println(m1)
	//print
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	//delete
	delete(m1, 1) // this will delete 1
	fmt.Println(m1)

	//clearing the map
	for k := range m1 {
		delete(m1, k)
	}

	fmt.Println(m1)

	//Exercise 1 -> Array
	x := [5]int{42, 43, 44, 45, 46}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)

	fmt.Println("----------------")

	//Exercise 2 -> Slice
	y := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range y {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", y)

	fmt.Println("----------------")

	//Exercise 3
	fmt.Println(y[:5])
	fmt.Println(y[5:]) //Başlangıç indexi dahil, bitiş indexi dahil değildir.
	fmt.Println(y[2:7])
	fmt.Println(y[1:6])
	fmt.Println(y)

	fmt.Println("----------------")

	//Exercise 4 -> Slice'a veri ekleme
	y = append(y, 52)
	fmt.Println(y)

	y = append(y, 53, 54, 55)
	fmt.Println(y)

	z := []int{56, 57, 58, 59, 60}

	y = append(y, z...)
	fmt.Println(y)

	fmt.Println("----------------")

	//Exercise 5 -> Slice'dan veri silme (Bunun için yine append kullanılacak)
	y1 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y1 = append(y1[:3], y1[6:]...)
	fmt.Println(y1)

	fmt.Println("----------------")

	//Exercise 6 -> Multi-dimentional slice (Slice of a slice)

	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooo, James."}
	fmt.Println(xs1, xs2)

	xxs := [][]string{xs1, xs2} //Multi dimentional slice
	fmt.Println(xxs)

	for i, xs := range xxs {
		fmt.Println("Record: ", i)
		for j, val := range xs {
			fmt.Printf("\t Index position: %v \t value: %v\n", j, val)
		}
	}

	fmt.Println("----------------")

	//Exercise 7 -> Map
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being Evil", "Ice cream", "Sunsets"},
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println("This is the record for ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

	fmt.Println("----------------")

	//Exercise 8 -> Add value to map

	m["fleming_ian"] = []string{"Steaks", "Cigars", "Espionage"}

	for k, v := range m {
		fmt.Println("This is the record for ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

	fmt.Println("----------------")

	//Exercise 9 -> Delete value from map

	delete(m, "no_dr")

	for k, v := range m {
		fmt.Println("This is the record for ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
