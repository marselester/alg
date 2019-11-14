package substr

// Bruteforce finds the first occurrence of a pattern string pat in a text string s.
// It requires ~m*n character compares to search for a pattern of length m in a text of length n, in the worst case.
//
// The function keeps one pointer (i) into the text and another pointer (j) into the pattern.
// For each i, it resets j to 0 and increments it until finding a mismatch
// or the end of the pattern (j == m).
// If we reach the end of the text (i == n-m+1) before the end of the pattern, then there is no match.
//
// In a typical text, j index rarely increments, so the running time is proportional to n.
// Nearly all of the compares find a mismatch with the first character of the pattern.
func Bruteforce(s, pat string) int {
	txtEnd := len(s) - len(pat) + 1

	for i := 0; i < txtEnd; i++ {
		j := 0
		for ; j < len(pat); j++ {
			if s[i+j] != pat[j] {
				break
			}
		}
		// Search hit.
		if j == len(pat) {
			return i
		}
	}
	// Search miss.
	return -1
}
