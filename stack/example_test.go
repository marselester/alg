package stack_test

import (
	"fmt"

	"github.com/marselester/alg/stack"
)

func ExampleArray() {
	s := stack.Array{}
	s.Push("fizz")
	s.Push("bazz")
	fmt.Printf("%q %q %q", s.Pop(), s.Pop(), s.Pop())
	// Output: "bazz" "fizz" ""
}

func ExampleLinkedList() {
	s := stack.LinkedList{}
	s.Push("fizz")
	s.Push("bazz")
	fmt.Printf("%q %q %q", s.Pop(), s.Pop(), s.Pop())
	// Output: "bazz" "fizz" ""
}
