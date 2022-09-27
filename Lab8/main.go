package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"os"
	"sort"
)

func main() {

	//Marshal komutunda kullanılan structın içerisindeki değişkenlerin isimleri büyük harfle başlamalıdır. Yoksa boş döner

	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs)) //Eğer struct içerisindeki değişkenlerin isimleri küçük harfle başlıyorsa boş döner -> [{},{}]

	fmt.Println("----------------")

	//Unmarshalling

	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs2 := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs2) //[]uint8 -> byte slice -> byte alias for uint8

	//peopleJson := []personJson{}
	var peopleJson []personJson

	errJson := json.Unmarshal(bs2, &peopleJson)
	if errJson != nil {
		fmt.Println(errJson)
	}
	fmt.Println("All of the data :", peopleJson)

	for i, v := range peopleJson {
		fmt.Println("Person number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

	fmt.Println("----------------")

	//Writer interface

	fmt.Println("Hello, playground")
	fmt.Fprintln(os.Stdout, "Hello, playground")
	io.WriteString(os.Stdout, "Hello, playground\n") //All of them is same

	fmt.Println("----------------")

	//Sort
	xi := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	xs := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println("Before sort")
	fmt.Println(xi)
	fmt.Println(xs)

	fmt.Println("After sort")
	sort.Ints(xi)
	fmt.Println(xi)

	sort.Strings(xs)
	fmt.Println(xs)

	fmt.Println("----------------")

	//Custom sort

	//Sort by age
	ps1 := personSort{"James", 32}
	ps2 := personSort{"Moneypenny", 27}
	ps3 := personSort{"M", 54}
	ps4 := personSort{"Q", 64}

	peopleSort := []personSort{ps1, ps2, ps3, ps4}
	fmt.Println(peopleSort)
	fmt.Println("Sorting by age")
	sort.Sort(ByAge(peopleSort))
	fmt.Println(peopleSort)
	fmt.Println("----------------")
	//Sort by name
	fmt.Println(peopleSort)
	fmt.Println("Sorting by name")
	sort.Sort(ByName(peopleSort))
	fmt.Println(peopleSort)
	fmt.Println("----------------")

	//bcrypt

	b := `password123`
	bs, err = bcrypt.GenerateFromPassword([]byte(b), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
	fmt.Println(bs)

	loginPassword1 := `password123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword1)) //Karşılaştırma yapılır.
	if err != nil {
		fmt.Println("You can't login")
		return
	}
	fmt.Println("You're logged in")

	fmt.Println("----------------")

	//Exercise 1

	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	bs, err = json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	fmt.Println("----------------")

	//Exercise 2
	ex2 := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(ex2)

	var ex2Json []ex2Json
	errEx2 := json.Unmarshal([]byte(ex2), &ex2Json)
	if errEx2 != nil {
		fmt.Println(errEx2)
	}
	fmt.Println(ex2Json)
	for i, v := range ex2Json {
		fmt.Println("Person number", i)
		fmt.Println("\t", v.First, v.Last, v.Age)
		for _, v2 := range v.Sayings {
			fmt.Println("\t\t", v2)
		}
	}

	fmt.Println("----------------")

	//Exercise 3

	errEx3 := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("We did something wrong", errEx3)
	}

	fmt.Println("----------------")

	//Exercise 4

	ex4i := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	ex4s := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(ex4i)
	// sort xi
	sort.Ints(ex4i)
	fmt.Println("After sort")
	fmt.Println(ex4i)
	fmt.Println("\t----------------")
	fmt.Println(ex4s)
	// sort xs
	sort.Strings(ex4s)
	fmt.Println("After sort")
	fmt.Println(ex4s)

	fmt.Println("----------------")

	//Exercise 5

	ex51 := ex5User{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	ex52 := ex5User{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	ex53 := ex5User{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	usersEx5 := []ex5User{ex51, ex52, ex53}
	fmt.Println("After sort by age")
	fmt.Println("----------------")
	sort.Sort(ByAgeEx5(usersEx5))
	for i, v := range usersEx5 {
		fmt.Println("Person number", i)
		fmt.Println("\t", v.First, v.Last, v.Age)
		for _, v2 := range v.Sayings {
			fmt.Println("\t\t", v2)
		}
	}

	fmt.Println("----------------")
	fmt.Println("After sort by last")
	fmt.Println("----------------")
	sort.Sort(ByLastEx5(usersEx5))
	for i, v := range usersEx5 {
		fmt.Println("Person number", i)
		fmt.Println("\t", v.First, v.Last, v.Age)
		for _, v2 := range v.Sayings {
			fmt.Println("\t\t", v2)
		}
	}

}

type person struct {
	First string
	Last  string
	Age   int
}

type personJson struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

type personSort struct {
	first string
	age   int
}

type ByAge []personSort

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

type ByName []personSort

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].first < a[j].first }

type user struct {
	First string
	Age   int
}

type ex2Json struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

type ex5User struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

type ByAgeEx5 []ex5User

func (a ByAgeEx5) Len() int           { return len(a) }
func (a ByAgeEx5) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAgeEx5) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLastEx5 []ex5User

func (a ByLastEx5) Len() int           { return len(a) }
func (a ByLastEx5) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLastEx5) Less(i, j int) bool { return a[i].Last < a[j].Last }
