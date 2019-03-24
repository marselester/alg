// Program fileindex takes file names from the command line and uses a symbol table
// to build an inverted index (values are used to locate keys) associating every word
// in any of the files with a set of file names where the word can be found,
// then takes keyword queries from stdin, and produces its associated list of files.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// set represents a set of file names.
type set map[string]struct{}

func main() {
	st := make(map[string]set)
	for _, filename := range os.Args[1:] {
		f, err := os.Open(filename)
		if err != nil {
			log.Fatalf("fileindex: %v", err)
		}

		err = buildInvertedIndex(f, st, filename)
		if err != nil {
			log.Fatalf("fileindex: failed to build index: %v", err)
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		var filenames []string
		for k := range st[word] {
			filenames = append(filenames, k)
		}
		fmt.Printf("%s: %s\n", word, filenames)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("fileindex: failed to scan words: %v", err)
	}
}

// buildInvertedIndex associates every word from input with a set of filenames.
func buildInvertedIndex(input io.Reader, st map[string]set, filename string) error {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		if st[word] == nil {
			st[word] = make(set)
		}
		st[word][filename] = struct{}{}
	}
	return scanner.Err()
}
