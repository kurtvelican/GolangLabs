package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	/*n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)*/

	//---------------------------------------------

	/*var answer1, answer2, answer3 string

	fmt.Println("Name :")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}
	fmt.Println("Fav Food :")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}
	fmt.Println("Fav Sport :")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}

	fmt.Println(answer1, answer2, answer3)*/

	//---------------------------------------------

	/*f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("James Bond")

	io.Copy(f, r)*/

	//---------------------------------------------

	/*f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))*/

	//log.Println is used to log errors. That means it will print the error to the console with log time

	//Exercise 1

	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln("JSON did not marshal - here is the error:", err)
	}
	fmt.Println(string(bs))

	fmt.Println("--------------------------  ")

	//Exercise 2

	bs, err = toJSON(p1)
	if err != nil {
		log.Fatalln(err) //we can use other log commands.
	}
	fmt.Println(string(bs))

	fmt.Println("--------------------")

	//Exercise 3

	c1 := customError{
		info: "Need more coffee",
	}

	foo(c1)
	fmt.Println("--------------------")

	//Exercise 4

	_, err = sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

type person struct {
	First   string
	Last    string
	Sayings []string
}

type customError struct {
	info string
}

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("Math error: %v %v %v", se.lat, se.long, se.err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//e := errors.New("more coffee needed")
		e := fmt.Errorf("more coffee needed - value was %v", f)
		return 0, sqrtError{"50.2289 N", "99.4656 W", e}
	}
	return 42, nil
}

func (ce customError) Error() string {
	return fmt.Sprintf("Here is the error: %v", ce.info)
}

func foo(e error) {
	//fmt.Println("foo ran -", e)
	fmt.Println("foo ran -", e, "\n", e.(customError).info)
}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("There was an error in toJSON: %v", err)
	}
	return bs, nil
}
