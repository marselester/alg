package spt

import "fmt"

func ExampleTopologicalSort() {
	g := NewAdjacencyList(13)
	g.Add(&Edge{0, 6, 0})
	g.Add(&Edge{0, 1, 0})
	g.Add(&Edge{0, 5, 0})
	g.Add(&Edge{5, 4, 0})
	g.Add(&Edge{3, 5, 0})
	g.Add(&Edge{2, 3, 0})
	g.Add(&Edge{2, 0, 0})
	g.Add(&Edge{6, 9, 0})
	g.Add(&Edge{6, 4, 0})
	g.Add(&Edge{8, 7, 0})
	g.Add(&Edge{7, 6, 0})
	g.Add(&Edge{9, 10, 0})
	g.Add(&Edge{9, 11, 0})
	g.Add(&Edge{9, 12, 0})
	g.Add(&Edge{11, 12, 0})

	/*
		The expected preorder is 8->7 2->3 0->6->9->10 11->12 1 5->4.
		Calculus -> Linear Algebra
		Introduction to CS -> Advanced Programming
		Algorithms -> Theoretical CS -> Artificial Intelligence -> Robotics
		Machine Learning -> Neural Networks
		Databases
		Scientific Computing -> Computational Biology
	*/
	fmt.Println(topologicalSort(g))
	// Output:
	// [8 7 2 3 0 6 9 10 11 12 1 5 4]
}
