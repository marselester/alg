package zerosum

import (
	"testing"
)

func TestPairCountNaive(t *testing.T) {
	var tt = []struct {
		nums []int
		want int
	}{
		{
			nums: nil,
			want: 0,
		},
		{
			nums: []int{324110, -442472, 626686, -324110},
			want: 1,
		},
	}
	for _, tc := range tt {
		got := PairCountNaive(tc.nums)
		if got != tc.want {
			t.Errorf("PairCountNaive(%v) = %d, want %d", tc.nums, got, tc.want)
		}
	}
}

func TestPairCountFast(t *testing.T) {
	var tt = []struct {
		nums []int
		want int
	}{
		{
			nums: nil,
			want: 0,
		},
		{
			nums: []int{324110, -442472, 626686, -324110},
			want: 1,
		},
	}
	for _, tc := range tt {
		got := PairCountFast(tc.nums)
		if got != tc.want {
			t.Errorf("PairCountFast(%v) = %d, want %d", tc.nums, got, tc.want)
		}
	}
}

func TestTripleCountNaive(t *testing.T) {
	var tt = []struct {
		nums []int
		want int
	}{
		{
			nums: nil,
			want: 0,
		},
		{
			nums: []int{324110, -442472, 626686, -157678, 508681, 123414, -77867, 155091, 129801, 287381},
			want: 1,
		},
	}
	for _, tc := range tt {
		got := TripleCountNaive(tc.nums)
		if got != tc.want {
			t.Errorf("TripleCountNaive(%v) = %d, want %d", tc.nums, got, tc.want)
		}
	}
}

func TestTripleCountFast(t *testing.T) {
	var tt = []struct {
		nums []int
		want int
	}{
		{
			nums: nil,
			want: 0,
		},
		{
			nums: []int{324110, -442472, 626686, -157678, 508681, 123414, -77867, 155091, 129801, 287381},
			want: 1,
		},
	}
	for _, tc := range tt {
		got := TripleCountFast(tc.nums)
		if got != tc.want {
			t.Errorf("TripleCountFast(%v) = %d, want %d", tc.nums, got, tc.want)
		}
	}
}
