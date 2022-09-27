package main

import "fmt"

func nestedFunc() int {
	fmt.Println(5)
	return 5
}

func nestedFunc2() int {
	return nestedFunc()
}
