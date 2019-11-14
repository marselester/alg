package substr

import "testing"

func TestBruteforce(t *testing.T) {
	tests := []struct {
		text    string
		pattern string
		want    int
	}{
		{"ABACADABRA", "ABRA", 6},
		{"ABACADABRA", "AB", 0},
		{"AAAAAAAAAB", "AAAAB", 5},
		{"日本語", "語", 6},
		{"日本語", "👩", -1},
	}

	for _, tc := range tests {
		got := Bruteforce(tc.text, tc.pattern)
		if got != tc.want {
			t.Errorf("Bruteforce(%q, %q) = %d, want %d", tc.text, tc.pattern, got, tc.want)
		}
	}
}
