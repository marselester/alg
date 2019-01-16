package quick

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

func TestPartition(t *testing.T) {
	tt := []struct {
		a    []string
		want []string
	}{
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
			a:    []string{"K", "R", "A", "T", "E", "L", "E", "P", "U", "I", "M", "Q", "C", "X", "O", "S"},
			want: []string{"E", "C", "A", "I", "E", "K", "L", "P", "U", "T", "M", "Q", "R", "X", "O", "S"},
		},
	}
	for _, tc := range tt {
		partition(tc.a, 0, len(tc.a)-1)
		if !equal(tc.a, tc.want) {
			t.Errorf("partition got %v, want %v", tc.a, tc.want)
		}
	}
}

func TestQuicksort(t *testing.T) {
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
			a:    []string{"Q", "U", "I", "C", "K", "S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"},
			want: []string{"A", "C", "E", "E", "I", "K", "L", "M", "O", "P", "Q", "R", "S", "T", "U", "X"},
		},
	}
	for _, tc := range tt {
		Sort(tc.a)
		if !equal(tc.a, tc.want) {
			t.Errorf("Sort() got %v, want %v", tc.a, tc.want)
		}
	}
}
