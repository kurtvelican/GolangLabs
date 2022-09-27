package main

import "fmt"

type calculator interface {
	sum(int1 int, int2 int)
	minus(int1 int, int2 int)
	multiply(int1 int, int2 int)
	divide(float1 float32, float2 float32)
	write(string1 string)
}

func sum(int1 int, int2 int) interface{} {
	fmt.Println(int1 + int2)
	return int1 + int2
}

func minus(int1 int, int2 int) interface{} {
	fmt.Println(int1 - int2)
	return int1 - int2
}

func multiply(int1 int, int2 int) interface{} {
	fmt.Println(int1 * int2)
	return int1 * int2
}

func divide(float1 float32, float2 float32) interface{} {
	fmt.Println(float1 / float2)
	return float1 / float2
}

func write(string1 string) interface{} {
	fmt.Println(string1)
	return string1
}
