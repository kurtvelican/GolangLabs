package Callback

func OddSumCallback(f func(xi ...int) int, vi ...int) int { //f func(xi ...int) kısmı callback'tir.
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
