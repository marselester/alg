package algebra_test

import (
	"fmt"

	algebra "github.com/marselester/alg/ebra"
)

func Example() {
	a := [][]float64{
		{0, .9, 0, 0, 0},
		{0, 0, .36, .36, .18},
		{0, 0, 0, .9, 0},
		{.9, 0, 0, 0, 0},
		{.47, 0, .47, 0, 0},
	}
	x := []float64{.05, .04, .36, .37, .19}
	b := make([]float64, len(x))

	var sum float64
	for i := 0; i < len(a); i++ {
		sum = 0
		for j := 0; j < len(a[i]); j++ {
			sum += a[i][j] * x[j]
		}
		b[i] = sum
	}
	fmt.Printf("%0.4f\n", b)

	m := make([]*algebra.SparseVector, len(a))
	for i := 0; i < len(a); i++ {
		m[i] = algebra.NewSparseVector(a[i])
	}
	for i := 0; i < len(a); i++ {
		b[i] = m[i].Dot(x)
	}
	fmt.Printf("%0.4f\n", b)
	// Output:
	// [0.0360 0.2970 0.3330 0.0450 0.1927]
	// [0.0360 0.2970 0.3330 0.0450 0.1927]
}
