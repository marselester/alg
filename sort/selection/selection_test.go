package selection

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

func TestSelectionSort(t *testing.T) {
	tt := []struct {
		a    []string
		want []string
	}{
		{a: nil, want: nil},
		{
			a:    []string{"B", "A"},
			want: []string{"A", "B"},
		},
		{
			a:    []string{"S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"},
			want: []string{"A", "E", "E", "L", "M", "O", "P", "R", "S", "T", "X"},
		},
	}
	for _, tc := range tt {
		Sort(tc.a)
		if !equal(tc.a, tc.want) {
			t.Errorf("Sort() got %v, want %v", tc.a, tc.want)
		}
	}
}
