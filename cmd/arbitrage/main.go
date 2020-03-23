/*
Program arbitrage finds an arbitrage opportunity in a currency exchange table by constructing
a complete-graph representation of the exchange table and then using Bellman-Ford algorithm
to find a negative cycle in the graph.

Consider a market for financial transactions that is based on trading commodities.
This table is equivalent to a complete edge-weighted digraph with a vertex corresponding to each currency
and an edge corresponding to each conversion rate.

	USD 1      0.741  0.657  1.061  1.005
	EUR 1.349  1      0.888  1.433  1.366
	GBP 1.521  1.126  1      1.614  1.538
	CHF 0.942  0.698  0.619  1      0.953
	CAD 0.995  0.732  0.650  1.049  1

One line per currency represents a currency name, and conversion rates to the other currencies.
For example, one USD buys 0.741 EUR, 0.657 GBP, 1.061 CHF, 1.005 CAD.
The line number (USD #0) indicates how much one unit (USD #0) buys units of currencies (EUR #1, GBP #2, CHF #3, CAD #4).

Another example, one EUR buys 1.3666 CAD, because
the line number (EUR #1) indicates how much one unit (EUR #1) buys units of CAD #4.

Paths in the graph specify multistep conversions, e.g., USD->EUR->CAD represents a way to convert one unit of USD into CAD.
For example, one USD buys 0.741 EUR which buys CAD with 1.366 per unit; that gives 0.741 * 1.366 = 1.012206 CAD.
That gives a better rate than directly converting USD to CAD: 1.012206 > 1.005.

Finding the path such that the product of the weights in maximal is certainly of interest (even better if it's > 1).
For example, if we convert 1.012206 CAD back to USD, we get 1.00714497 USD: 0.741 * 1.366 * 0.995 = 1.00714497 USD.
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/marselester/alg/digraph/spt"
	"github.com/marselester/alg/digraph/weighted"
)

func main() {
	var currenciesCount int
	flag.IntVar(&currenciesCount, "currencies", 0, "number of currencies (vertices) in the digraph")
	var stake float64
	flag.Float64Var(&stake, "stake", 1000.0, "stake")
	flag.Parse()

	currencies, edges, err := findOpportunity(currenciesCount, os.Stdin)
	if err != nil {
		log.Fatalf("arbitrage: %v", err)
	}

	for _, e := range edges {
		fmt.Printf("%0.5f %s ", stake, currencies[e.V])
		stake *= math.Exp(-e.Weight)
		fmt.Printf("= %0.5f %s\n", stake, currencies[e.W])
	}
}

func findOpportunity(currenciesCount int, input io.Reader) ([]string, []*weighted.Edge, error) {
	if currenciesCount <= 0 {
		return nil, nil, fmt.Errorf("number of currencies must be positive: %d", currenciesCount)
	}

	dag := weighted.NewAdjacencyList(currenciesCount)
	currencies := make([]string, currenciesCount)
	curIndex := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		s := scanner.Text()
		name, rates, err := parseRates(s)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to parse %q: %w", s, err)
		}
		currencies[curIndex] = name

		for curPairIndex, r := range rates {
			dag.Add(&weighted.Edge{
				V:      curIndex,
				W:      curPairIndex,
				Weight: -math.Log(r),
			})
		}

		curIndex++
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("failed to read a digraph: %w", err)
	}

	sp := spt.NewBellmanFord(dag, 0)
	if sp.HasNegativeCycle() {
		edges := sp.NegativeCycle()
		return currencies, edges, nil
	}
	return nil, nil, nil
}

func parseRates(s string) (currency string, rates []float64, err error) {
	ss := strings.Fields(s)
	if len(ss) == 0 {
		return "", nil, fmt.Errorf("empty currency rates")
	}

	currency = ss[0]
	// Currency name should be in ISO 4217. For simplicity's sake only first character is checked.
	okName := 'A' <= currency[0] && currency[0] <= 'Z'
	if !okName {
		return "", nil, fmt.Errorf("invalid currency: %v", currency)
	}

	// There are no rates provided.
	if len(ss) == 1 {
		return currency, nil, fmt.Errorf("no rates")
	}

	rates = make([]float64, 0, len(ss)-1)
	for i := 1; i < len(ss); i++ {
		r, err := strconv.ParseFloat(ss[i], 64)
		if err != nil {
			return "", nil, fmt.Errorf("invalid rate format: %w", err)
		}
		rates = append(rates, r)
	}

	return currency, rates, nil
}
