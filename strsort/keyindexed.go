package strsort

import (
	"reflect"
)

/*
KeyIndexedCounting is a stable sorting method that is effective (liner-time sort)
whenever the keys are small integers (radix is within a constant factor of the slice length).

The key function should return a small number to sort the slice by.
It panics if the provided interface is not a slice.
Radix is an alphabet size (number of characters), for example, binary is 2, DNA is 4, octal is 8.

The first step is to count the frequency of occurence of each key.
Next, compute the starting index positions in the sorted order of items.
In general, to get the starting index for items with any given key we sum the frequency counts of smaller values.
For example, there are three items with key 1 and five items with key 2,
then the items with key 3 start at position 8 in the sorted array.

With the count array transformed into an index table, we accomplish the actual sort by moving the items
to an auxiliary array aux. We move each item to the position in aux indicated by the count entry
corresponding to its key, and then increment that entry.
For each key i, count[i] is the index of the position in aux where the next item with key i should be placed.
*/
func KeyIndexedCounting(slice interface{}, key func(i int) int, radix int) {
	rv := reflect.ValueOf(slice)
	// Note, count[0] is always zero.
	count := make([]int, radix+1)
	for i := 0; i < rv.Len(); i++ {
		// Store frequency of the key in the next element.
		count[key(i)+1]++
	}

	// Transform counts to indices. For example, at first []count represents key frequences [0 0 3 5 6 6].
	// Then it represents starting index positions [0 0 3 8 14 20].
	for r := 0; r < radix; r++ {
		count[r+1] += count[r]
	}

	// Distribute the records.
	aux := reflect.MakeSlice(rv.Type(), rv.Len(), rv.Len())
	var pos int
	for i := 0; i < rv.Len(); i++ {
		pos = count[key(i)]
		aux.Index(pos).Set(rv.Index(i))
		count[key(i)]++
	}

	// Copy back.
	for i := 0; i < rv.Len(); i++ {
		rv.Index(i).Set(aux.Index(i))
	}
}
