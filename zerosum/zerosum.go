// Package zerosum counts the number of unordered pairs (2-sum),
// triples (3-sum) of n distinct integers that sum to zero.
package zerosum

import (
	"sort"

	"github.com/marselester/alg/search/binary-search"
)

// PairCountNaive implements the brute-force 2-sum algorithm that has
// quadratic order of growth of the running time.
func PairCountNaive(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == 0 {
				count++
			}
		}
	}
	return count
}

// PairCountFast implements 2-sum algorithm with linearithmic (n log n) order of growth
// of the running time. The algorithm is based on the fact that an entry
// nums[i] is one of a pair that sums to zero if the value -nums[i] is in the array.
// Firstly, the array is sorted (to enable binary search), then each item
// is binary searched for its negative value.
func PairCountFast(nums []int) int {
	sort.Ints(nums)
	count := 0
	for i := 0; i < len(nums); i++ {
		foundUniq := bsearch.IndexOf(nums, -nums[i]) > i
		if foundUniq {
			count++
		}
	}
	return count
}

// TripleCountNaive implements the brute-force 3-sum algorithm that has
// cubic order of growth of the running time.
func TripleCountNaive(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					count++
				}
			}
		}
	}
	return count
}

// TripleCountFast implements 3-sum algorithm with n² log n order of growth
// of the running time. The same idea from 2-sum problem is effective for 3-sum.
// A pair of nums[i] + nums[j] is part of a triple that sums to zero
// if the value -(nums[i] + nums[j]) is in the array.
// Firstly, the array is sorted to enable binary search.
func TripleCountFast(nums []int) int {
	sort.Ints(nums)
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			foundUniq := bsearch.IndexOf(nums, -(nums[i]+nums[j])) > j
			if foundUniq {
				count++
			}
		}
	}
	return count
}
