// Package bsearch implements binary search algorithm.
// Its running time's order of growth is logarithmic.
package bsearch

// IndexOf looks up a search key in the sorted array nums and
// returns the index of the key if it is present in the array,
// -1 otherwise.
func IndexOf(nums []int, key int) int {
	var lo, hi, mid int
	hi = len(nums) - 1

	for hi >= lo {
		mid = lo + (hi-lo)/2
		switch {
		case key > nums[mid]:
			lo = mid + 1
		case key < nums[mid]:
			hi = mid - 1
		default:
			return mid
		}
	}

	return -1
}
