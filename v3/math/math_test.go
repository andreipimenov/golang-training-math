package math_test

import (
	"testing"

	"github.com/andreipimenov/golang-training-math/v3/math"
)

func TestMultiply(t *testing.T) {
	result := math.Multiply(2, 5)
	expected := 10
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"normal", []int{1, 2, 3, -5}, 1},
		{"empty", []int{}, 0},
		{"one entry", []int{99}, 99},
		{"all zeros", []int{0, 0, 0, 0, 0, 0, 0, 0, 0}, 0},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			result := math.Add(test.input...)
			if test.expected != result {
				t.Fatalf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestPow(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		n        int
		expected float64
	}{
		{"2^5", 2, 5, 32},
		{"2^0", 10, 0, 1},
		{"2^-1", 2, -1, 0.5},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			// t.Log("Running ", test.n)
			result := math.Pow(test.input, test.n)
			if test.expected != result {
				t.Fatalf("expected %v, got %v", test.expected, result)
			}
			// t.Fail()
		})
	}
}
