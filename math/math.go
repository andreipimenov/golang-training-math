package math

func Sum(v ...int) int {
	r := 0
	for _, i := range v {
		r += i
	}
	return r
}

func Multiply(a, b int) int {
	return a * b
}
