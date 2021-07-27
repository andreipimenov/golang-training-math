package math_test

import (
	stdmath "math"
	"testing"

	"github.com/andreipimenov/golang-training-math/v3/math"
)

func BenchmarkPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Pow(i, i)
	}
}

func BenchmarkStdPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stdmath.Pow(float64(i), float64(i))
	}
}
