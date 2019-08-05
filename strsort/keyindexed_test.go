package strsort

import (
	"fmt"
	"testing"
)

func ExampleKeyIndexedCounting() {
	// There are 5 groups numbered in range [0; 4], though nobody is in group zero.
	const groups = 5
	students := []student{
		{"Anderson", 2},
		{"Brown", 3},
		{"Davis", 3},
		{"Garsia", 4},
		{"Harris", 1},
		{"Jackson", 3},
		{"Johnson", 4},
		{"Jones", 3},
		{"Martin", 1},
		{"Martinez", 2},
		{"Miller", 2},
		{"Moore", 1},
		{"Robinson", 2},
		{"Smith", 4},
		{"Taylor", 3},
		{"Thomas", 4},
		{"Thompson", 4},
		{"White", 2},
		{"Williams", 3},
		{"Wilson", 4},
	}

	KeyIndexedCounting(
		students,
		func(i int) int { return students[i].group },
		groups,
	)
	for _, s := range students {
		fmt.Println(s.name, s.group)
	}
	// Output:
	// Harris 1
	// Martin 1
	// Moore 1
	// Anderson 2
	// Martinez 2
	// Miller 2
	// Robinson 2
	// White 2
	// Brown 3
	// Davis 3
	// Jackson 3
	// Jones 3
	// Taylor 3
	// Williams 3
	// Garsia 4
	// Johnson 4
	// Smith 4
	// Thomas 4
	// Thompson 4
	// Wilson 4
}

func TestKeyIndexedCounting(t *testing.T) {
	const groups = 5
	students := []student{
		{"Anderson", 2},
		{"Brown", 3},
		{"Davis", 3},
		{"Garsia", 4},
		{"Harris", 1},
		{"Jackson", 3},
		{"Johnson", 4},
		{"Jones", 3},
		{"Martin", 1},
		{"Martinez", 2},
		{"Miller", 2},
		{"Moore", 1},
		{"Robinson", 2},
		{"Smith", 4},
		{"Taylor", 3},
		{"Thomas", 4},
		{"Thompson", 4},
		{"White", 2},
		{"Williams", 3},
		{"Wilson", 4},
	}
	want := []student{
		{"Harris", 1},
		{"Martin", 1},
		{"Moore", 1},
		{"Anderson", 2},
		{"Martinez", 2},
		{"Miller", 2},
		{"Robinson", 2},
		{"White", 2},
		{"Brown", 3},
		{"Davis", 3},
		{"Jackson", 3},
		{"Jones", 3},
		{"Taylor", 3},
		{"Williams", 3},
		{"Garsia", 4},
		{"Johnson", 4},
		{"Smith", 4},
		{"Thomas", 4},
		{"Thompson", 4},
		{"Wilson", 4},
	}

	keyIndexedCounting(students, groups)
	for i, s := range want {
		if students[i] != s {
			t.Fatalf("keyIndexedCounting() got %v, want %v", students, want)
		}
	}
}
