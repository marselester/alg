package zerosum

import (
	"math/rand"
	"testing"
	"time"
)

func randNums(n int) []int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = r.Int()
		// Flip a sign of an odd number.
		if nums[i]&1 == 1 {
			nums[i] *= -1
		}
	}
	return nums
}

func BenchmarkPairCountNaive(b *testing.B) {
	var tt = []struct {
		name string
		nums []int
	}{
		{"size=250", randNums(250)},
		{"size=500", randNums(500)},
		{"size=1K", randNums(1e3)},
		{"size=2K", randNums(2e3)},
		{"size=4K", randNums(4e3)},
	}
	b.ResetTimer()

	for _, tc := range tt {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				PairCountNaive(tc.nums)
			}
		})
	}
}

func BenchmarkPairCountFast(b *testing.B) {
	var tt = []struct {
		name string
		nums []int
	}{
		{"size=250", randNums(250)},
		{"size=500", randNums(500)},
		{"size=1K", randNums(1e3)},
		{"size=2K", randNums(2e3)},
		{"size=4K", randNums(4e3)},
	}
	b.ResetTimer()

	for _, tc := range tt {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				PairCountFast(tc.nums)
			}
		})
	}
}

func BenchmarkTripleCountNaive(b *testing.B) {
	var tt = []struct {
		name string
		nums []int
	}{
		{"size=250", randNums(250)},
		{"size=500", randNums(500)},
		{"size=1K", randNums(1e3)},
		{"size=2K", randNums(2e3)},
		{"size=4K", randNums(4e3)},
	}
	b.ResetTimer()

	for _, tc := range tt {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				TripleCountNaive(tc.nums)
			}
		})
	}
}

func BenchmarkTripleCountFast(b *testing.B) {
	var tt = []struct {
		name string
		nums []int
	}{
		{"size=250", randNums(250)},
		{"size=500", randNums(500)},
		{"size=1K", randNums(1e3)},
		{"size=2K", randNums(2e3)},
		{"size=4K", randNums(4e3)},
	}
	b.ResetTimer()

	for _, tc := range tt {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				TripleCountFast(tc.nums)
			}
		})
	}
}
