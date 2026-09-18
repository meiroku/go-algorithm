package prime

import (
	"slices"
	"testing"
)

func TestSieveOfEratosthenes(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected []int
	}{
		{name: "2未満(負)", input: -5, expected: nil},
		{name: "2未満(0)", input: 0, expected: nil},
		{name: "2未満(1)", input: 1, expected: nil},
		{name: "最小の素数(2)", input: 2, expected: []int{2}},
		{name: "10以下の素数", input: 10, expected: []int{2, 3, 5, 7}},
		{name: "30以下の素数", input: 30, expected: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SieveOfEratosthenes(tt.input)
			if !slices.Equal(got, tt.expected) {
				t.Errorf("SieveOfEratosthenes(%d) = %v; want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func BenchmarkSieveOfEratosthenes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SieveOfEratosthenes(1000000)
	}
}