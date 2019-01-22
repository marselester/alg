package qfind

import (
	"math/rand"
	"testing"
	"time"
)

func equal(seq1, seq2 [][2]int) bool {
	if len(seq1) != len(seq2) {
		return false
	}
	for i := 0; i < len(seq1); i++ {
		if seq1[i] != seq2[i] {
			return false
		}
	}
	return true
}

func TestNetwork(t *testing.T) {
	net := New(10)
	seq := [][2]int{
		{4, 3},
		{3, 8},
		{6, 5},
		{9, 4},
		{2, 1},
		{8, 9},
		{5, 0},
		{7, 2},
		{6, 1},
		{1, 0},
		{6, 7},
	}
	want := [][2]int{
		{4, 3},
		{3, 8},
		{6, 5},
		{9, 4},
		{2, 1},
		// {8, 9},
		{5, 0},
		{7, 2},
		{6, 1},
		// {1, 0},
		// {6, 7},
	}
	var got [][2]int

	for _, pair := range seq {
		if !net.IsConnected(pair[0], pair[1]) {
			net.Connect(pair[0], pair[1])
			got = append(got, pair)
		}
	}
	if !equal(got, want) {
		t.Errorf("Network quick-find connectivity is %v, want %v", got, want)
	}
}

func randPairs(n int) [][2]int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	pairs := make([][2]int, n)
	for i := 0; i < n; i++ {
		pairs[i] = [2]int{
			r.Int() % n,
			r.Int() % n,
		}
	}
	return pairs
}

func BenchmarkQuickFind(b *testing.B) {
	var tt = []struct {
		name  string
		net   *Network
		pairs [][2]int
	}{
		{"size=1K", New(1e3), randPairs(1e3)},
		{"size=10K", New(1e4), randPairs(1e4)},
		{"size=100K", New(1e5), randPairs(1e5)},
		{"size=200K", New(2e5), randPairs(2e5)},
	}
	b.ResetTimer()

	for _, tc := range tt {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, pair := range tc.pairs {
					tc.net.Connect(pair[0], pair[1])
				}
			}
		})
	}
}
