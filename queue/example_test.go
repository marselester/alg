package queue_test

import (
	"fmt"

	"github.com/marselester/alg/queue"
)

func ExampleArray() {
	s := queue.Array{}
	s.Enqueue("fizz")
	s.Enqueue("bazz")
	fmt.Printf("%q %q %q", s.Dequeue(), s.Dequeue(), s.Dequeue())
	// Output: "fizz" "bazz" ""
}

func ExampleLinkedList() {
	q := queue.LinkedList{}
	q.Enqueue("fizz")
	q.Enqueue("bazz")
	fmt.Printf("%q %q %q", q.Dequeue(), q.Dequeue(), q.Dequeue())
	// Output: "fizz" "bazz" ""
}

func ExampleRingBuffer() {
	rb := queue.NewRingBuffer(1)
	fmt.Println(rb.Enqueue("fizz"))
	fmt.Println(rb.Enqueue("bazz"))
	fmt.Println(rb.Dequeue())
	// Output:
	// true
	// false
	// fizz true
}
