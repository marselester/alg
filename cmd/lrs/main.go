// Command lrs finds the longest repeated substring in the text on standard input
// by building a suffix array and then scanning through the sorted suffixes to find
// the longest value.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"github.com/marselester/alg/suffixarray"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("lrs: failed to read from standard input: %v", err)
	}

	re := regexp.MustCompile(`\s+`)
	text := re.ReplaceAllString(string(b), " ")
	fmt.Printf("%q\n", suffixarray.LongestRepeatedSubstring(text))
}
