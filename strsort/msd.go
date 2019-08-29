package strsort

const (
	// DefaultCutoff is a cutoff for small subarrays threshold (insertion sort).
	DefaultCutoff = 15
	// DefaultRadix is extended ASCII alphabet size (number of characters).
	DefaultRadix = 256
)

type msd struct {
	cutoff int // cutoff for small subarrays.
	radix  int
	aux    []string
}
type msdOption func(*msd)

// WithCutoff defines a size of a subarray when to use insertion sort algorithm.
func WithCutoff(cutoff int) msdOption {
	return func(s *msd) {
		s.cutoff = cutoff
	}
}

// WithRadix defines alphabet size (number of characters).
func WithRadix(radix int) msdOption {
	return func(s *msd) {
		s.radix = radix
	}
}

/*
MSD performs most-significant-digit first (MSD) string sort where strings are not necessarily all the same length.
It uses key-indexed counting to sort the strings according to their first character,
then recursively sorts the subarrays corresponding to each character
(excluding the first character which is the same).

The cost of MSD string sort depends strongly on the number of possible characters.
To sort n random strings from R-character alphabet, MSD sort examines n*log n
characters (logarithm's base is R) on average.
In the best case MSD uses just one pass.
The worst case for MSD is when all strings are equal (linear running time like LSD string sort).

The amount of space needed is proportional to R times the length of the longest string (plus n)
in the worst case.

The main challenge in getting max efficiency from MSD sort on long strings is to deal with lack of randomness.
*/
func MSD(a []string, options ...msdOption) {
	s := msd{
		cutoff: DefaultCutoff,
		radix:  DefaultRadix,
		aux:    make([]string, len(a)),
	}
	for _, opt := range options {
		opt(&s)
	}

	s.sort(a, 0, len(a)-1, 0)
}

// sort sorts a slice of strings on their first character using key-indexed counting,
// then recursively sorts the subarrays corresponding to each first-character value.
func (s *msd) sort(a []string, lo, hi, column int) {
	if hi <= lo+s.cutoff {
		insertionSort(a, lo, hi)
		return
	}

	// Compute frequency counts.
	count := make([]int, s.radix+2)
	for i := lo; i <= hi; i++ {
		count[charAt(a[i], column)+2]++
	}

	// Transform counts to indices.
	for r := 0; r < s.radix+1; r++ {
		count[r+1] += count[r]
	}

	// Distribute.
	for i := lo; i <= hi; i++ {
		c := charAt(a[i], column) + 1
		s.aux[count[c]] = a[i]
		count[c]++
	}

	// Copy back.
	for i := lo; i <= hi; i++ {
		a[i] = s.aux[i-lo]
	}

	// Recursively sort for each character value excluding -1.
	for r := 0; r < s.radix; r++ {
		s.sort(a, lo+count[r], lo+count[r+1]-1, column+1)
	}
}

/*
charAt converts from an indexed string character to an array index that returns -1
if the specified character position is past the end of the string.
This convention means that we have radix+1 different possible character values at each string position:

	-1 to signify end of string
	0 for the first alphabet character
	1 for the second alphabet character, and so forth

When using charAt return value, add 1 to get a nonnegative int that can be used to index arrays.
*/
func charAt(str string, column int) int {
	if column < len(str) {
		return int(str[column])
	}
	return -1
}

// insertionSort sorts from a[lo] to a[hi] in increasing order using insertion sort algorithm.
func insertionSort(a []string, lo, hi int) {
	for i := lo; i <= hi; i++ {
		for j := i; j > lo; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}
