// Program parentheses reads in a text stream from standard input
// and uses a stack to determine whether its parentheses are properly balanced.
// For example, [()]{}{[()()]()} is a balanced, whereas [(]) is not.
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/marselester/alg/stack"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var expr string
	for scanner.Scan() {
		expr = scanner.Text()
		fmt.Println(isBalanced(expr))
	}
}

func isBalanced(expr string) bool {
	s := stack.LinkedList{}
	for _, r := range expr {
		switch par := string(r); par {
		case "[", "(", "{":
			s.Push(par)
		case "]", ")", "}":
			if ok := bracketsMatch(s.Pop(), par); !ok {
				return false
			}
		}
	}

	allMatched := s.Size() == 0
	return allMatched
}

func bracketsMatch(openpar, closepar string) bool {
	switch openpar {
	case "[":
		return closepar == "]"
	case "(":
		return closepar == ")"
	case "{":
		return closepar == "}"
	}
	return false
}
