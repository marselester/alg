/*
Package selection implements selection sort algorithm.
First, it finds the smallest item in the array and exchanges it with the first entry.
Then, it finds the next smallest item and exchanges it with the second entry.

Selection sort uses n²/2 compares and n exchanges to sort an array of length n.
The number of exchanges is a linear function of the array length.

Running time is insensitive to initial order in the input array: it takes about as
long to sort a randomly-ordered array and array that is already ordered.
*/
package selection

// Sort sorts a slice of strings in increasing order by repeatedly selecting
// the smallest remaining item.
func Sort(a []string) {
	var min int
	for i := 0; i < len(a); i++ {
		// min is an index of min value found in the slice.
		min = i

		// The entries to the left of position i are the i smallest items in the array
		// and are not examined again.
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}

		// Exchange min item with the current entry.
		a[i], a[min] = a[min], a[i]
	}
}
