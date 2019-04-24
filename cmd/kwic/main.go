// Program kwic is a keyword-in-context indexing client.
// It builds a suffix array for the text from standard input, takes a query from a flag,
// and prints search results.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"github.com/marselester/alg/suffixarray"
)

func main() {
	query := flag.String("query", "", "keyword to look for in standard input")
	ctxlen := flag.Int("len", 15, "number of characters after query to give context")
	flag.Parse()
	if len(*query) == 0 {
		log.Fatal("kwic: search query must not be blank")
	}

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("kwic: failed to read from standard input: %v", err)
	}

	re := regexp.MustCompile(`\s+`)
	text := re.ReplaceAllString(string(b), " ")
	idx := suffixarray.New(text)
	for _, s := range idx.Search(*query, *ctxlen) {
		fmt.Println(s)
	}
}
