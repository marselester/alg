/*
Program whitelist is a demo for binary search algorithm.
Imagine a credit card company that needs to check whether customer transactions are
for a valid account. To do so, it can
- keep customers account numbers in a whitelist file,
- produce the account number associated with each transaction in the standard input,
- put onto standard output the numbers that are not associated with any customer.
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

	"github.com/marselester/alg/search/binary-search"
)

func main() {
	whitelistPath := flag.String("whitelist", "", "path to a whitelist file")
	txPath := flag.String("tx", "", "optional path to a transactions file")
	flag.Parse()

	var err error
	wfile, err := os.Open(*whitelistPath)
	if err != nil {
		log.Fatalf("whitelist: failed to open the whitelist: %v", err)
	}
	defer wfile.Close()

	tfile := os.Stdin
	if *txPath != "" {
		tfile, err = os.Open(*txPath)
		if err != nil {
			log.Fatalf("whitelist: failed to open the transactions: %v", err)
		}
		defer tfile.Close()
	}

	whitelist, err := loadAllAccounts(wfile)
	if err != nil {
		log.Fatalf("whitelist: failed to load the whitelist: %v", err)
	}
	accounts, err := loadAllAccounts(tfile)
	if err != nil {
		log.Fatalf("whitelist: failed to load the transactions: %v", err)
	}

	for acc := range accounts {
		if i := bsearch.IndexOf(whitelist, acc); i == -1 {
			fmt.Println(acc)
		}
	}
}

func loadAllAccounts(in io.Reader) (accounts []int, err error) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)

	var i int
	for scanner.Scan() {
		if i, err = strconv.Atoi(scanner.Text()); err != nil {
			return
		}
		accounts = append(accounts, i)
	}
	err = scanner.Err()
	return
}
