// Program multiway (multiway merge priority-queue) merges together the sorted input streams
// into a single sorted output stream.
// Each stream index is associated with an item (the next char in the stream).
// It prints in a loop the smallest char in the queue and removes the corresponding entry,
// then adds a new entry for the next char in that stream.
//
// Usage example:
//     $ ./multiway ABCFGIIZ BDHPQQ ABEFJN
//     AABBBCDEFFGHIIJNPQQZ
//
// The streams might be the outputs of scientific instruments (sorted by time),
// commercial transactions (sorted by account number or time).
// You can read input streams and put them in sorted order on the output
// no matter how long they are.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/marselester/alg/sort/pqueue"
)

func main() {
	streams := make([]*bufio.Reader, len(os.Args)-1)
	for i, s := range os.Args[1:] {
		streams[i] = bufio.NewReader(strings.NewReader(s))
	}

	// n is number of streams.
	n := len(streams)
	pq := pqueue.NewIndexMinHeap(n)

	for i := 0; i < n; i++ {
		char, _, err := streams[i].ReadRune()
		if err != nil {
			if err == io.EOF {
				continue
			}
			log.Fatalf("multiway: failed to read a char: %v", err)
		}

		item := float64(char)
		pq.Insert(i, item)
	}

	for pq.Size() != 0 {
		i, item := pq.Min()
		fmt.Printf("%c", rune(item))

		char, _, err := streams[i].ReadRune()
		if err != nil {
			if err == io.EOF {
				continue
			}
			log.Fatalf("multiway: failed to read a char: %v", err)
		}

		item = float64(char)
		pq.Insert(i, item)
	}
	fmt.Println()
}
