package merge

import "testing"

func equal(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestMerge(t *testing.T) {
	tt := []struct {
		a    []string
		lo   int
		mid  int
		hi   int
		want []string
	}{
		{a: nil, want: nil},
		{
			a:    []string{},
			want: []string{},
		},
		{
			a:    []string{"A"},
			want: []string{"A"},
		},
		{
			a:    []string{"B", "A"},
			mid:  0,
			hi:   1,
			want: []string{"A", "B"},
		},
		{
			a:    []string{"E", "E", "G", "M", "R", "A", "C", "E", "R", "T"},
			mid:  4,
			hi:   9,
			want: []string{"A", "C", "E", "E", "E", "G", "M", "R", "R", "T"},
		},
	}
	for _, tc := range tt {
		ms := &mergesort{
			a:   tc.a,
			aux: make([]string, len(tc.a)),
		}
		ms.merge(tc.lo, tc.mid, tc.hi)
		if !equal(tc.a, tc.want) {
			t.Errorf("merge(%d, %d, %d) got %v, want %v", tc.lo, tc.mid, tc.hi, tc.a, tc.want)
		}
	}
}

func TestTopdownMergesort(t *testing.T) {
	tt := []struct {
		a    []string
		want []string
	}{
		{a: nil, want: nil},
		{
			a:    []string{},
			want: []string{},
		},
		{
			a:    []string{"A"},
			want: []string{"A"},
		},
		{
			a:    []string{"B", "A"},
			want: []string{"A", "B"},
		},
		{
			a:    []string{"B", "C", "A"},
			want: []string{"A", "B", "C"},
		},
		{
			a:    []string{"S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"},
			want: []string{"A", "E", "E", "L", "M", "O", "P", "R", "S", "T", "X"},
		},
		{
			a:    []string{"M", "E", "R", "G", "E", "S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"},
			want: []string{"A", "E", "E", "E", "E", "G", "L", "M", "M", "O", "P", "R", "R", "S", "T", "X"},
		},
	}
	for _, tc := range tt {
		TDSort(tc.a)
		if !equal(tc.a, tc.want) {
			t.Errorf("TDSort() got %v, want %v", tc.a, tc.want)
		}
	}
}

func TestBottomupMergesort(t *testing.T) {
	tt := []struct {
		a    []string
		want []string
	}{
		{a: nil, want: nil},
		{
			a:    []string{},
			want: []string{},
		},
		{
			a:    []string{"A"},
			want: []string{"A"},
		},
		{
			a:    []string{"B", "A"},
			want: []string{"A", "B"},
		},
		{
			a:    []string{"B", "C", "A"},
			want: []string{"A", "B", "C"},
		},
		{
			a:    []string{"S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"},
			want: []string{"A", "E", "E", "L", "M", "O", "P", "R", "S", "T", "X"},
		},
		{
			a:    []string{"M", "E", "R", "G", "E", "S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"},
			want: []string{"A", "E", "E", "E", "E", "G", "L", "M", "M", "O", "P", "R", "R", "S", "T", "X"},
		},
	}
	for _, tc := range tt {
		BUSort(tc.a)
		if !equal(tc.a, tc.want) {
			t.Errorf("BUSort() got %v, want %v", tc.a, tc.want)
		}
	}
}
