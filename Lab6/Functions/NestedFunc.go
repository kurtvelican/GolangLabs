package Functions

import "fmt"

func NestedFunc() int {
	fmt.Println(5)
	return 5
}

func NestedFunc2() int {
	return NestedFunc()
}
