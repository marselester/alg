package shell

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

func TestShellSort(t *testing.T) {
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
			a:    []string{"S", "H", "E", "L", "L", "S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"},
			want: []string{"A", "E", "E", "E", "H", "L", "L", "L", "M", "O", "P", "R", "S", "S", "T", "X"},
		},
	}
	for _, tc := range tt {
		Sort(tc.a)
		if !equal(tc.a, tc.want) {
			t.Errorf("Sort() got %v, want %v", tc.a, tc.want)
		}
	}
}
