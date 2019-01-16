/*
Package insertion provides insertion sort algorithm that people often use to sort
playing cards: take one card at a time from a deck, inserting each into its proper place
among those already in hand. We need to make space to insert the current item by
moving larger items one position to the right, before inserting the current item into
the vacated position.

We take the first card from a deck, then one more (index 1).
Then we go left from the current card position and compare it with "known" cards in hand.
If the card is larger than the current card, then we slide it to the right in one position.
If not, then we're done sliding.

It  works well for partially sorted arrays and a fine method for tiny arrays.
The worst case is n²/2 compares and n²/2 exchanges and the best case is n-1 compares and
zero exchanges.

https://www.khanacademy.org/computing/computer-science/algorithms/insertion-sort/a/insertion-sort
*/
package insertion

// Sort sorts a slice of strings in increasing order using insertion sort algorithm.
func Sort(a []string) {
	// Each new position i is like the new card handed to you by the dealer.
	for i := 1; i < len(a); i++ {
		// You need to insert the new card into the correct place in
		// the sorted subarray to the left of that position.
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}
