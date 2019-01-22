/*
Command 3sum counts the number of unordered triples in a file of n distinct integers that sum to zero.
For example, second, eighth, and tenth entries sum to 0:
324110, -442472, 626686, -157678, 508681, 123414, -77867, 155091, 129801, 287381.

The program is useful to study relationship between the problem size n and running time of 3sum.
You can download data sets from https://algs4.cs.princeton.edu/code/algs4-data.zip.
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/marselester/alg/zerosum"
)

func main() {
	filename := flag.String("file", "", "path to a file where ints are stored")
	fast := flag.Bool("fast", false, "enable fast 3-sum algorithm")
	flag.Parse()

	var err error
	f := os.Stdin
	if *filename != "" {
		f, err = os.Open(*filename)
		if err != nil {
			log.Fatalf("3sum: failed to open %q: %v", *filename, err)
		}
		defer f.Close()
	}

	nums, err := readAllInts(f)
	if err != nil {
		log.Fatalf("3sum: failed to read all ints: %v", err)
	}
	if *fast {
		fmt.Println(zerosum.TripleCountFast(nums))
	} else {
		fmt.Println(zerosum.TripleCountNaive(nums))
	}
}

func readAllInts(in io.Reader) (nums []int, err error) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)

	var i int
	for scanner.Scan() {
		if i, err = strconv.Atoi(scanner.Text()); err != nil {
			return
		}
		nums = append(nums, i)
	}
	err = scanner.Err()
	return
}
