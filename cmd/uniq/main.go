// Program uniq counts the distinct words in the input using quicksort.
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/marselester/alg/sort/quick"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	quick.Sort(words)
	count := 1
	for i := 1; i < len(words); i++ {
		if words[i] != words[i-1] {
			count++
		}
	}
	fmt.Println(count)
}
