package strsort

/*
LSD performs least-significant-digit first (LSD) string sort given that the strings are fixed-length,
for example, license plates, telephone numbers, bank account numbers, IP addresses.

If the strings are each of length w, we sort the strings w times with key-indexed counting,
using each of the positions as the key, proceeding from right to left.

LSD is a linear-time stable string sort. Total running time is proportional to w*n.
For typical applications radix is far smaller than slice length n.
No matter how large n, it makes w passes through the data.

	~ 10*w*n + n + 4*w*radix array accesses
	extra space proportional to n + radix

*/
func LSD(a []string, radix int) {
	if len(a) == 0 {
		return
	}
	// For simplicity's sake assume all the strings are the same length and each byte represents one character.
	w := len(a[0])
	aux := make([]string, len(a))

	for column := w - 1; column >= 0; column-- {
		count := make([]int, radix+1)
		for i := range a {
			count[a[i][column]+1]++
		}

		for r := 0; r < radix; r++ {
			count[r+1] += count[r]
		}

		for i := range a {
			pos := count[a[i][column]]
			aux[pos] = a[i]
			count[a[i][column]]++
		}

		for i := range a {
			a[i] = aux[i]
		}
	}
}
