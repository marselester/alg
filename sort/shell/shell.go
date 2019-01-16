/*
Package shell implements shellsort algorithm which is an extension of insertion sort
that gains speed by allowing entries exchanges that are far apart, to produce
partially sorted arrays that can be efficiently sorted by insertion sort.

Insertion sort is slow for large unordered arrays because it only exchanges
adjacent entries. For example, if the smallest item is located at the end
of the array, then n-1 exchanges are needed to move it to the start.

This implementation uses 3*h + 1 increment sequence for steps.
The worst-case number of compares is n**1.5 (not quadratic).
*/
package shell

// Sort sorts a slice of strings in increasing order using shellsort algorithm.
func Sort(a []string) {
	// Start sorting from short subsequences (large steps) and ending at step h=1.
	for h := step(len(a)); h >= 1; h = h / 3 {
		hsort(a, h)
	}
}

// hsort takes items from the slice with step h and sorts them in increasing order.
// For example, when step is 4, every 4th item is compared and swapped if necessary.
func hsort(a []string, h int) {
	for i := h; i < len(a); i++ {
		// Insertion sort is modified to decrement by step h, instead of 1.
		for j := i; j >= h; j = j - h {
			if a[j] < a[j-h] {
				a[j], a[j-h] = a[j-h], a[j]
			}
		}
	}
}

// step calculates the largest step size to start moving items in long distances.
func step(n int) int {
	var h int
	for h = 1; h < n/3; h = 3*h + 1 {
	}
	return h
}
