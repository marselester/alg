package digraph

import "fmt"

func ExampleTopologicalSort() {
	g := NewAdjacencyList(13)
	g.Add(0, 6)
	g.Add(0, 1)
	g.Add(0, 5)
	g.Add(5, 4)
	g.Add(3, 5)
	g.Add(2, 3)
	g.Add(2, 0)
	g.Add(6, 9)
	g.Add(6, 4)
	g.Add(8, 7)
	g.Add(7, 6)
	g.Add(9, 10)
	g.Add(9, 11)
	g.Add(9, 12)
	g.Add(11, 12)

	/*
		The expected preorder is 8->7 2->3 0->6->9->10 11->12 1 5->4.
		Calculus -> Linear Algebra
		Introduction to CS -> Advanced Programming
		Algorithms -> Theoretical CS -> Artificial Intelligence -> Robotics
		Machine Learning -> Neural Networks
		Databases
		Scientific Computing -> Computational Biology
	*/
	fmt.Println(TopologicalSort(g))
	// Output:
	// [8 7 2 3 0 6 9 10 11 12 1 5 4]
}
