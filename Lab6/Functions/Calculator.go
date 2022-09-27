package Functions

import "fmt"

type Calculator interface {
	Sum(int1 int, int2 int)
	Minus(int1 int, int2 int)
	Multiply(int1 int, int2 int)
	Divide(float1 float32, float2 float32)
	Write(string1 string)
}

func Sum(int1 int, int2 int) interface{} {
	fmt.Println(int1 + int2)
	return int1 + int2
}

func Minus(int1 int, int2 int) interface{} {
	fmt.Println(int1 - int2)
	return int1 - int2
}

func Multiply(int1 int, int2 int) interface{} {
	fmt.Println(int1 * int2)
	return int1 * int2
}

func Divide(float1 float32, float2 float32) interface{} {
	fmt.Println(float1 / float2)
	return float1 / float2
}

func Write(string1 string) interface{} {
	fmt.Println(string1)
	return string1
}
