// Program toptx finds top n transactions in the huge input stream (considered unbounded so it can't be sorted)
// using priority queue. Comparing each new transaction against n largest seen so far is likely
// to be expensive unless n is small.
//
// Expected output is 4747.08 4732.35 4409.74 4381.21 4121.85 of the following input:
// 644.08 4121.85 2678.40 4409.74 837.42 3229.27 4732.35 4381.21 66.10 4747.08 2156.86 1025.70 2520.97 708.95 3532.36 4050.20.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/marselester/alg/sort/pqueue"
	"github.com/marselester/alg/stack"
)

func main() {
	n := flag.Int("n", 5, "Number of top transactions to show.")
	flag.Parse()

	pq := pqueue.NewMinHeap(*n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var amount float64
	var err error
	for scanner.Scan() {
		if amount, err = strconv.ParseFloat(scanner.Text(), 64); err != nil {
			log.Fatalf("toptx: failed to parse amount: %v", err)
		}
		pq.Insert(amount)
		if pq.Size() > *n {
			pq.Min()
		}
	}

	s := stack.NewArray(*n)
	for pq.Size() > 0 {
		s.Push(fmt.Sprintf("%v", pq.Min()))
	}
	for s.Size() > 0 {
		fmt.Println(s.Pop())
	}
}
