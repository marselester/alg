package main

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func ExampleFindOpportunity() {
	table := strings.NewReader(`USD 1      0.741  0.657  1.061  1.005
EUR 1.349  1      0.888  1.433  1.366
GBP 1.521  1.126  1      1.614  1.538
CHF 0.942  0.698  0.619  1      0.953
CAD 0.995  0.732  0.650  1.049  1`)
	currencies, edges, _ := findOpportunity(5, table)

	stake := 1000.0
	for _, e := range edges {
		fmt.Printf("%0.5f %s ", stake, currencies[e.V])
		stake *= math.Exp(-e.Weight)
		fmt.Printf("= %0.5f %s\n", stake, currencies[e.W])
	}
	// Output:
	// 1000.00000 USD = 741.00000 EUR
	// 741.00000 EUR = 1012.20600 CAD
	// 1012.20600 CAD = 1007.14497 USD
}

func TestParseRates(t *testing.T) {
	tests := map[string]struct {
		input    string
		currency string
		rates    []float64
	}{
		"USD": {
			"USD 1      0.741  0.657  1.061  1.005",
			"USD",
			[]float64{1, 0.741, 0.657, 1.061, 1.005},
		},
		"EUR": {
			"EUR 1.349  1      0.888  1.433  1.366",
			"EUR",
			[]float64{1.349, 1, 0.888, 1.433, 1.366},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotCurrency, gotRates, err := parseRates(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if tc.currency != gotCurrency {
				t.Fatalf("expected currency %v, got %v", tc.currency, gotCurrency)
			}
			if !equalFloat(tc.rates, gotRates) {
				t.Fatalf("expected rates %v, got %v", tc.rates, gotRates)
			}
		})
	}
}

func TestParseRatesError(t *testing.T) {
	tests := map[string]struct {
		input string
		err   string
	}{
		"blank":             {"", "empty currency rates"},
		"whitespace":        {" ", "empty currency rates"},
		"tab":               {"\t", "empty currency rates"},
		"three whitespaces": {"   ", "empty currency rates"},
		"blank name":        {" 0.1", "invalid currency: 0.1"},
		"currency only":     {"USD", "no rates"},
		"rate letter":       {"USD a", "invalid rate format: strconv.ParseFloat: parsing \"a\": invalid syntax"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, _, got := parseRates(tc.input)
			if got == nil || tc.err != got.Error() {
				t.Fatalf("expected %q, got %q", tc.err, got)
			}
		})
	}
}

func equalFloat(s1, s2 []float64) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
