package main

import (
	"fmt"
	"testing"
)

func equal(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestTokenize(t *testing.T) {
	tt := []struct {
		expr string
		want []string
	}{
		{"", nil},
		{"(1+2)", []string{"(", "1", "+", "2", ")"}},
		{"(1 + 2)", []string{"(", "1", "+", "2", ")"}},
		{"( 1 + 2 )", []string{"(", "1", "+", "2", ")"}},
		{"(    1  +   2 )    ", []string{"(", "1", "+", "2", ")"}},
	}

	for _, tc := range tt {
		t.Run(tc.expr, func(t *testing.T) {
			got := tokenize(tc.expr)
			if !equal(got, tc.want) {
				t.Errorf("tokenize(%q) = %q, want %q", tc.expr, got, tc.want)
			}
		})
	}
}

func TestEval(t *testing.T) {
	tt := []struct {
		op   string
		a    string
		b    string
		want string
	}{
		{op: "+", a: "1", b: "2", want: "3"},
		{op: "-", a: "3", b: "4", want: "-1"},
		{op: "*", a: "3", b: "4", want: "12"},
		{op: "/", a: "3", b: "4", want: "0"},
		{op: "/", a: "6", b: "2", want: "3"},
	}

	for _, tc := range tt {
		expr := fmt.Sprintf("%s %s %s", tc.a, tc.op, tc.b)
		t.Run(expr, func(t *testing.T) {
			got, err := eval(tc.op, tc.a, tc.b)
			if err != nil {
				t.Fatal(err)
			}
			if got != tc.want {
				t.Errorf("eval(%q, %q, %q) = %q, want %q", tc.op, tc.a, tc.b, got, tc.want)
			}
		})
	}
}
