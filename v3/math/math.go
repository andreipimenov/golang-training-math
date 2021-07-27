package math

func Add(v ...int) int {
	r := 0
	for _, i := range v {
		r += i
	}
	return r
}

func Multiply(v ...int) int {
	r := 1
	for _, i := range v {
		r *= i
	}
	return r
}

func Pow(v, n int) float64 {

	// time.Sleep(1 * time.Second)

	p := n
	if n < 1 {
		p = -n
	}
	r := 1.0
	for i := 0; i < p; i++ {
		r *= float64(v)
	}
	if n < 1 {
		r = 1 / r
	}
	return r
}
