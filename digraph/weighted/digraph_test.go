package weighted

import (
	"fmt"
	"testing"
)

func TestPathTo(t *testing.T) {
	edgeTo := []*Edge{
		nil,
		{5, 1, 0.32},
		{0, 2, 0.26},
		{7, 3, 0.39},
		{0, 4, 0.38},
		{4, 5, 0.35},
		{3, 6, 0.52},
		{2, 7, 0.34},
	}
	want := []*Edge{
		{0, 2, 0.26},
		{2, 7, 0.34},
		{7, 3, 0.39},
		{3, 6, 0.52},
	}

	got := PathTo(6, edgeTo)
	if fmt.Sprint(want) != fmt.Sprint(got) {
		t.Errorf("PathTo(6, %v) = %v want %v", edgeTo, got, want)
	}
}
