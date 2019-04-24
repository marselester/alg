/*
Package suffixarray provides substring search using suffix array — the abstraction of a sorted list of
suffix strings. The cost of comparing two suffixes may be proportional to the length of the suffixes
in the case when their common prefix is very long, but most comparisons in typical applications involve
only a few characters. If so, the running time of the suffix sort is linearithmic.
If all the characters are equal, the sort examines every character in each suffix and thus takes
quadratic time.
*/
package suffixarray

import (
	"sort"
)

// LongestPrefix returns a longest common prefix of given two strings.
// It takes time proportional to the length of the match.
func LongestPrefix(s1, s2 string) string {
	// n is a length of the shortest string.
	var n int
	if len(s1) < len(s2) {
		n = len(s1)
	} else {
		n = len(s2)
	}

	for i := 0; i < n; i++ {
		if s1[i] != s2[i] {
			n = i
			break
		}
	}

	return s1[:n]
}

/*
LongestRepeatedSubstring returns the longest substring that appears at least twice in a given string.
The idea is to make an array of the n suffixes (the substrings starting at each position and
going to the end), and then sort this array. Every substring appears somewhere as a prefix of one
of the suffixes in the array. After sorting, the longest repeated substrings will appear in adjacent
positions in the array. Thus, we can make a single pass through the sorted array, keeping track of
the longest matching prefixes between adjacent strings.
*/
func LongestRepeatedSubstring(s string) string {
	if len(s) < 2 {
		return ""
	}

	suffix := make([]string, len(s))
	// For simplicity's sake, s is considered to be ASCII string.
	for i := 0; i < len(s); i++ {
		suffix[i] = s[i:]
	}
	sort.Strings(suffix)

	var longest string
	for i, j := 0, 1; j < len(suffix); i, j = i+1, j+1 {
		p := LongestPrefix(suffix[i], suffix[j])
		if len(p) > len(longest) {
			longest = p
		}
	}
	return longest
}

// New returns a new KeywordIndex index to find a substring in a text.
// For simplicity's sake, a text is considered to be ASCII string.
func New(text string) *KeywordIndex {
	ki := KeywordIndex{
		suffix: make([]string, len(text)),
	}
	for i := 0; i < len(text); i++ {
		ki.suffix[i] = text[i:]
	}
	sort.Strings(ki.suffix)
	return &ki
}

/*
KeywordIndex represents a sorted list of suffix strings to find a substring within a large text.
The idea is to make an array of the n suffixes (the substrings starting at each position and
going to the end), and then sort this array. A binary search is used to search in that array,
comparing the search key with each suffix.
*/
type KeywordIndex struct {
	suffix []string
}

// Search returns all occurrences that matched the query including specified number of characters
// after to give context.
func (ki *KeywordIndex) Search(query string, ctxlen int) []string {
	var found []string
	for i := ki.rank(query); i < len(ki.suffix); i++ {
		if p := LongestPrefix(query, ki.suffix[i]); len(query) == len(p) {
			n := min(len(ki.suffix[i]), len(query)+ctxlen)
			found = append(found, ki.suffix[i][:n])
		}
	}
	return found
}

// rank returns the number of suffixes less than the given search key.
// It helps to find the first possible suffix in the sorted suffix list that has key as prefix
// and that all other occurrences of key in the text immediately follow.
// If a prefix is found, its array index is returned (index equals to number of keys smaller than the search key).
// Otherwise, it also returns the number of keys that are smaller than the search key.
func (ki *KeywordIndex) rank(s string) int {
	var lo, mid, hi int
	hi = len(ki.suffix) - 1

	for lo <= hi {
		mid = lo + (hi-lo)/2
		switch {
		case s == ki.suffix[mid]:
			return mid
		case s > ki.suffix[mid]:
			lo = mid + 1
		case s < ki.suffix[mid]:
			hi = mid - 1
		}
	}

	// lo always equals to number of keys that are smaller than the search key.
	return lo
}
func min(i, j int) int {
	if i >= j {
		return j
	}
	return i
}
