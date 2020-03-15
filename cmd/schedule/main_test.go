package main

import (
	"strings"
	"testing"
)

func TestFindSchedule(t *testing.T) {
	spec := strings.NewReader(`41.0 1 7 9
51.0 2
50.0
36.0
38.0
45.0
21.0 3 8
32.0 3 8
32.0 2
29.0 4 6`)
	minTime, startTimes, err := findSchedule(10, spec)
	if err != nil {
		t.Fatal(err)
	}

	wantMinTime := 173.0
	if wantMinTime != minTime {
		t.Errorf("expected min time %0.2f, got %0.2f", wantMinTime, minTime)
	}

	wantStartTimes := []float64{
		0,
		41,
		123,
		91,
		70,
		0,
		70,
		41,
		91,
		41,
	}
	if !equalFloat(wantStartTimes, startTimes) {
		t.Fatalf("expected start times %v, got %v", wantStartTimes, startTimes)
	}
}

func TestParseJobConstraint(t *testing.T) {
	tests := map[string]struct {
		input    string
		duration float64
		jobs     []int
	}{
		"three jobs":         {"41.1 1 7 9", 41.1, []int{1, 7, 9}},
		"one job":            {"51.5 2", 51.5, []int{2}},
		"no jobs":            {"50", 50, nil},
		"negative duration":  {"-50", -50, nil},
		"negative job index": {"50 -1", 50, []int{-1}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotDuration, gotJobs, err := parseJobConstraint(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if tc.duration != gotDuration {
				t.Fatalf("expected duration %v, got %v", tc.duration, gotDuration)
			}
			if !equal(tc.jobs, gotJobs) {
				t.Fatalf("expected jobs %v, got %v", tc.jobs, gotJobs)
			}
		})
	}
}

func TestParseJobConstraintError(t *testing.T) {
	tests := map[string]struct {
		input string
		err   string
	}{
		"blank":             {"", "empty job"},
		"whitespace":        {" ", "empty job"},
		"tab":               {"\t", "empty job"},
		"three whitespaces": {"   ", "empty job"},
		"duration letter":   {"a", "invalid duration format: strconv.ParseFloat: parsing \"a\": invalid syntax"},
		"job letter":        {"1 a", "invalid job index format: strconv.Atoi: parsing \"a\": invalid syntax"},
		"job emoji":         {"50 😫", "invalid job index format: strconv.Atoi: parsing \"😫\": invalid syntax"},
		"both emojis":       {"😫 😫", "invalid duration format: strconv.ParseFloat: parsing \"😫\": invalid syntax"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, _, got := parseJobConstraint(tc.input)
			if tc.err != got.Error() {
				t.Fatalf("expected %q, got %q", tc.err, got)
			}
		})
	}
}

func equal(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func equalFloat(s1, s2 []float64) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
