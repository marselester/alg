package strsort

/*
Quick sorts a slice of strings using three-way string quicksort algorithm.
It 3-way partitions strings on their first character,
then recursively sorts three resulting subarrays (the strings whose first character is):

	less than the partitioning character
	equal to the partitioning character
	greater than the partitioning character

The essential idea behind 3-way quicksort is to take special action when the leading characters are equal.
In the small subarrays, where most of the compares in the sort are done, the strings are likely to have
numerous equal leading characters. The standard algorithm has to scan over all those characters for each compare;
the 3-way algorithm avoids doing so.

To sort an array of n random strings, it uses ~ 2 * n * ln n character compares on the average.
The algorithm has no direct dependencies on the size of the alphabet.

The algorithm outperforms other sorting methods when strings have long common prefixes that
it doesn't have to reexamine (web logs analysis).
*/
func Quick(a []string) {
	quick3way(a, 0, len(a)-1, 0)
}

func quick3way(a []string, lo, hi, column int) {
	if hi <= lo {
		return
	}

	lt, gt := lo, hi
	v := charAt(a[lo], column)

	for i := lo + 1; i <= gt; {
		switch t := charAt(a[i], column); {
		case t < v:
			a[lt], a[i] = a[i], a[lt]
			lt++
			i++
		case t > v:
			a[i], a[gt] = a[gt], a[i]
			gt--
		default:
			i++
		}
	}

	quick3way(a, lo, lt-1, column)
	if v >= 0 {
		quick3way(a, lt, gt, column+1)
	}
	quick3way(a, gt+1, hi, column)
}
