package suffixarray

import (
	"testing"
)

func TestLongestPrefix(t *testing.T) {
	tt := []struct {
		s1   string
		s2   string
		want string
	}{
		{"", "", ""},
		{"a", "b", ""},
		{"a", "a", "a"},
		{"acctgttaac", "accgttaa", "acc"},
	}

	for _, tc := range tt {
		got := LongestPrefix(tc.s1, tc.s2)
		if got != tc.want {
			t.Errorf("LongestPrefix(%q, %q) = %q, want %q", tc.s1, tc.s2, got, tc.want)
		}
	}
}

func TestLongestRepeatedSubstring(t *testing.T) {
	s := "AACAAGTTTACAAGC"
	want := "ACAAG"
	got := LongestRepeatedSubstring(s)
	if got != want {
		t.Errorf("LongestRepeatedSubstring(%q) = %q, want %q", s, got, want)
	}
}

func TestKeywordIndexRank(t *testing.T) {
	idx := New("it was the best of times it was the")
	s := "th"
	want := 30
	got := idx.rank(s)
	if got != want {
		t.Errorf("rank(%q) = %d, want %d", s, got, want)
	}
}

func TestKeywordIndexSearch(t *testing.T) {
	idx := New("it was the best of times it was the")
	s := "th"
	want := []string{"the", "the best of times it was the"}
	got := idx.Search(s, 100)
	if !equal(got, want) {
		t.Errorf("Search(%q) = %q, want %q", s, got, want)
	}
}

func BenchmarkLongestRepeatedSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LongestRepeatedSubstring("AACAAGTTTACAAGC")
	}
}

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
