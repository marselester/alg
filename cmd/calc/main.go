/*
Program calc computes the value of arithmetic expression like this one:

	(1 + ((2 + 3) * (4 * 5)))

It takes a string as input expression and produces the number represented by the expression as output.
An expression consists of parenthesis, operators, and operands (numbers).
The Dijkstra's two-stack arithmetic expression-evaluation algorithm:

- push operands onto the operand stack
- push operators onto the operator stack
- ignore left parentheses
- on encountering a right parenthesis, pop an operator, pop the requisite number of operands,
  and push onto the operand stack the result of applying that operator to those operands.

After the final right parenthesis has been processed, there is one value on the stack,
which is the value of the expression.
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/marselester/alg/stack"
)

func main() {
	expr := flag.String("expr", "", "Arithmetic expression to evaluate, e.g., (1 + ((2 + 3) * (4 * 5))).")
	flag.Parse()

	operand := stack.Array{}
	operator := stack.Array{}
	for _, t := range tokenize(*expr) {
		switch t {
		// Ignore left parentheses.
		case "(":
		// Push operators onto the operator stack.
		case "+", "-", "*", "/":
			operator.Push(t)
		// Evaluate the operation from the stack.
		case ")":
			op, b, a := operator.Pop(), operand.Pop(), operand.Pop()
			res, err := eval(op, a, b)
			if err != nil {
				log.Fatalf("calc: eval %s %s %s: %v", a, op, b, err)
			}
			operand.Push(res)
			fmt.Printf("%s %s %s = %s\n", a, op, b, res)
		// Push operands onto the operand stack.
		default:
			operand.Push(t)
		}
	}
}

// tokenize parses arithmetic expression and returns tokens to process.
func tokenize(expr string) []string {
	// Add whitespaces near parenthesis and operators to facilitate tokenization.
	// Whitespace is used as a word delimeter.
	ops := []string{"(", ")", "+", "-", "*", "/"}
	for _, op := range ops {
		expr = strings.Replace(expr, op, fmt.Sprintf(" %s ", op), -1)
	}

	var tt []string
	scanner := bufio.NewScanner(strings.NewReader(expr))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		tt = append(tt, scanner.Text())
	}
	return tt
}

// eval computes arithmetic expressions, e.g., a + b.
func eval(op, a, b string) (string, error) {
	x, err := strconv.Atoi(a)
	if err != nil {
		return "", err
	}
	y, err := strconv.Atoi(b)
	if err != nil {
		return "", err
	}

	switch op {
	case "+":
		return fmt.Sprintf("%d", x+y), nil
	case "-":
		return fmt.Sprintf("%d", x-y), nil
	case "*":
		return fmt.Sprintf("%d", x*y), nil
	case "/":
		if y == 0 {
			return "", fmt.Errorf("division by zero")
		}
		return fmt.Sprintf("%d", x/y), nil
	default:
		return "", fmt.Errorf("unknown operation")
	}
}
