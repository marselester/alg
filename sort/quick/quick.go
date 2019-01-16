/*
Package quick provides quicksort algorithm that is divide-and-conquer method for sorting.
It is in-place, requires time proportional to n log n (linearithmic).

Before sorting, the array is shuffled to eliminate dependence on input.
Then it's recursively partitioned. The partitioning method is as follows:
choose partitioning item (e.g., a[lo]),
scan from left to right until you find item >= partitioning item,
scan from right to left until you find item <= partitioning item.
These two items that stopped iterations are out of place, so we exchange them.
When the scan indices cross, exchange the partitioning item a[lo]
with the rightmost entry of the left subarray a[j] and return index j.

Quicksort might have quadratic running time if partitions are unbalanced.
For example, 1st partition is on the smallest item, 2nd -- on the next smallest
item and so forth, so the program removes just one item for each call, leading
to excessive number of partitions of large subarrays. Shuffling makes it unlikely
that bad partitions will happen consistently.

Quicksort is slower than insertion sort for tiny subarrays (5-15 items).
*/
package quick

import (
	"math/rand"
	"time"
)

// Sort sorts array in increasing order using quicksort algorithm.
func Sort(a []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	sort(a, 0, len(a)-1)
}

func sort(a []string, lo, hi int) {
	if lo >= hi {
		return
	}

	j := partition(a, lo, hi)
	sort(a, lo, j-1) // Sort left part.
	sort(a, j+1, hi) // Sort right part.
}

func partition(a []string, lo, hi int) int {
	i, j := lo, hi
	v := a[lo] // Partitioning item.

Partitioning:
	for {
		for ; a[i] <= v; i++ {
			if i == hi {
				break Partitioning
			}
		}
		for ; a[j] >= v; j-- {
			if j == lo {
				break Partitioning
			}
		}
		if i >= j {
			break Partitioning
		}
		a[i], a[j] = a[j], a[i]
	}

	a[lo], a[j] = a[j], a[lo]
	return j
}
