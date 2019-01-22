package bsearch

import "testing"

func TestIndexOf(t *testing.T) {
	var tt = []struct {
		nums []int
		key  int
		want int
	}{
		{
			nums: nil,
			key:  123,
			want: -1,
		},
		{
			nums: []int{123},
			key:  123,
			want: 0,
		},
		{
			nums: []int{1, 2},
			key:  123,
			want: -1,
		},
		{
			nums: []int{1, 2},
			key:  2,
			want: 1,
		},
		{
			nums: []int{10, 11, 12, 16, 18, 23, 29, 33, 48, 54, 57, 68, 77, 84, 98},
			key:  23,
			want: 5,
		},
		{
			nums: []int{10, 11, 12, 16, 18, 23, 29, 33, 48, 54, 57, 68, 77, 84, 98},
			key:  50,
			want: -1,
		},
	}
	for _, tc := range tt {
		got := IndexOf(tc.nums, tc.key)
		if got != tc.want {
			t.Errorf("IndexOf(%v, %d) = %d, want %d", tc.nums, tc.key, got, tc.want)
		}
	}
}
