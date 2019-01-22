package main

import "testing"

func TestIsBalanced(t *testing.T) {
	var tt = []struct {
		expr string
		want bool
	}{
		{"[", false},
		{"{", false},
		{"[", false},
		{"[)", false},
		{"[]]", false},
		{"[(])", false},
		{"[]", true},
		{"()", true},
		{"{}", true},
		{"[()]{}{[()()]()}", true},
	}

	for _, tc := range tt {
		t.Run(tc.expr, func(t *testing.T) {
			got := isBalanced(tc.expr)
			if got != tc.want {
				t.Errorf("isBalanced(%q) = %v, want %v", tc.expr, got, tc.want)
			}
		})
	}
}
