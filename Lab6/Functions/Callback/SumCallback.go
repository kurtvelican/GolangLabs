package Callback

func SumCallback(xi ...int) int {
	//fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
