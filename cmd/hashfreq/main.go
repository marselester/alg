// Program hashfreq calculates hash value frequencies for given words
// to ensure hash function spreads a typical set of keys uniformly among the values
// between 0 and m-1.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/marselester/alg/search/hashtable"
)

func main() {
	tablesize := flag.Int("m", 97, "hash table size")
	flag.Parse()

	// freq counts how many times the same hash value was produced.
	freq := make([]int, *tablesize)
	var hashval int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		hashval = hashtable.Hash(scanner.Text(), *tablesize)
		freq[hashval]++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("hashfreq: %v", err)
	}

	for k, v := range freq {
		fmt.Printf("%d, %d\n", k, v)
	}
}
