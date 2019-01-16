/*
Package merge implements mergesort algorithm. An array is divided in half,
two halves are sorted (recursively), and then results are merged.
Mergesort's running time is linearithmic (n log n). Its prime disadvantage is that
it uses extra space proportional to n.
*/
package merge

type mergesort struct {
	a   []string
	aux []string
}

// TDSort sorts array in increasing order using recursive top-down mergesort implementation
// (divide-and-conquer paradigm).
func TDSort(a []string) {
	ms := &mergesort{
		a:   a,
		aux: make([]string, len(a)),
	}
	ms.topdown(0, len(a)-1)
}

// BUSort sorts array in increasing order using bottom-up mergesort implementation:
// firstly it merges subarrays containing only one item,
// then it merges subarrays with two elements and so on, doubling the step on each pass.
func BUSort(a []string) {
	ms := &mergesort{
		a:   a,
		aux: make([]string, len(a)),
	}

	length := len(ms.a)
	for step := 1; step < length; step = step * 2 {
		for lo := 0; lo < length-step; lo = lo + step*2 {
			mid := lo + step - 1
			hi := mid + step
			if hi > length-1 {
				hi = length - 1
			}
			ms.merge(lo, mid, hi)
		}
	}
}

// topdown recursively sorts two halves and then merges them.
func (ms *mergesort) topdown(lo, hi int) {
	if lo >= hi {
		return
	}

	mid := lo + (hi-lo)/2
	// Sort left half.
	ms.topdown(lo, mid)
	// Sort right half.
	ms.topdown(mid+1, hi)
	// Merge two parts.
	ms.merge(lo, mid, hi)
}

// merge merges two halves a[lo:mid] and a[mid+1:hi] of the array in increasing order.
// It requires to copy the original array into auxiliary array.
func (ms *mergesort) merge(lo, mid, hi int) {
	if len(ms.a) == 0 {
		return
	}
	for k := lo; k <= hi; k++ {
		ms.aux[k] = ms.a[k]
	}

	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {
		switch {
		// Left half is exhausted (take from the right). Move right cursor one step.
		case i > mid:
			ms.a[k] = ms.aux[j]
			j++
		// Right half is exhausted (take from the left). Move left cursor one step.
		case j > hi:
			ms.a[k] = ms.aux[i]
			i++
		// Item on the right is less than on the left side (take from the right).
		// Move right cursor one step.
		case ms.aux[i] > ms.aux[j]:
			ms.a[k] = ms.aux[j]
			j++
		// Item on the left is less or equal to the right one (take from the left).
		// Move left cursor one step.
		case ms.aux[i] <= ms.aux[j]:
			ms.a[k] = ms.aux[i]
			i++
		}
	}
}
