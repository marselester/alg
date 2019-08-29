package strsort

import "fmt"

func ExampleMSD() {
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
	MSD(a, WithCutoff(0))

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
