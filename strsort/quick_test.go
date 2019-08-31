package strsort

import "fmt"

func ExampleQuick() {
	a := []string{
		"she",
		"sells",
		"seashells",
		"by",
		"the",
		"seashore",
		"the",
		"shells",
		"she",
		"sells",
		"are",
		"surely",
		"seashells",
	}
	Quick(a)

	for i := range a {
		fmt.Println(a[i])
	}
	// Output:
	// are
	// by
	// seashells
	// seashells
	// seashore
	// sells
	// sells
	// she
	// she
	// shells
	// surely
	// the
	// the
}
