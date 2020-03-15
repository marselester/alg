/*
Program schedule prints start times for each job using critical path method (linear-time).

The critical path method for parallel job scheduling:

	- create an edge-weighted DAG with a source, a sink, and two vertices for each job (start and end vertices)
	- for each job, add an edge from its start vertex to its end vertex with weight equal to its duration
	- for each precedence constraint v->w, add a zero-weight edge from the end vertex corresponding to v
	  to the beginning vertex corresponding to w
	- add zero-weight edges from the source to each job's start vertex
	- add zero-weight edges from each job's end vertex to the sink
	- schedule each job at the time given by the length of its longest path from the source

Parallel precedence-constrained scheduling is equivalent to a longest-paths problem in DAG.
We focus on the earliest possible time that we can schedule each job.
Here is an example input of ten jobs and their constraints:

	41.0 1 7 9
	51.0 2
	50.0
	36.0
	38.0
	45.0
	21.0 3 8
	32.0 3 8
	32.0 2
	29.0 4 6

Each line corresponds to a job index, e.g., first line describes job #0:
41.0 is a job duration, 1 7 9 are job indexes that must be completed before this job.

Duration of 173 is the minimum possible completion time for any schedule for this problem:
the schedule satisfies all the constraints, and no schedule can complete before 173
because of the job sequince 0->9->6->8->2 (this critical path is the earliest possible completion time).
It's called a critical path, because any delay in the starting time of any job
delays the best achievable completion time of the entire project.
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/marselester/alg/digraph/lpt"
	"github.com/marselester/alg/digraph/weighted"
)

func main() {
	var jobsCount int
	flag.IntVar(&jobsCount, "jobs", 0, "number of jobs (vertices) in the digraph")
	flag.Parse()

	minTime, startTimes, err := findSchedule(jobsCount, os.Stdin)
	if err != nil {
		log.Fatalf("schedule: %v", err)
	}

	fmt.Printf("Project finish time: %0.2f. Job start times:\n", minTime)
	for i, t := range startTimes {
		fmt.Printf("#%d: %0.2f\n", i, t)
	}
}

func findSchedule(jobsCount int, input io.Reader) (float64, []float64, error) {
	if jobsCount <= 0 {
		return -1, nil, fmt.Errorf("number of jobs must be positive: %d", jobsCount)
	}

	var (
		// Digraph will have 2*n + 2 vertices, because each job is represented as two vertices,
		// also there are source and sink vertices.
		verticesCount = 2*jobsCount + 2
		sourceIndex   = 2 * jobsCount
		sinkIndex     = sourceIndex + 1
	)
	dag := weighted.NewAdjacencyList(verticesCount)

	var (
		// First line in the input corresponds to job #0. It's also a "job start" vertex.
		jobStartIndex = 0
		// Each job has "job finish" vertex. For example, when there are 10 jobs,
		// 0..9 vertices represent job starts, 10..19 vertices represent job finishes,
		// 20 is the source vertex, and 21 is the sink.
		jobFinishIndex = jobStartIndex + jobsCount
	)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		s := scanner.Text()
		duration, jobConstraints, err := parseJobConstraint(s)
		if err != nil {
			return -1, nil, fmt.Errorf("unable to parse %q: %w", s, err)
		}
		if duration <= 0 {
			return -1, nil, fmt.Errorf("unable to parse %q: duration must be positive", s)
		}

		// Edge from the job's start vertex to the job's finish vertex
		// with weight equal to duration.
		dag.Add(&weighted.Edge{
			V:      jobStartIndex,
			W:      jobFinishIndex,
			Weight: duration,
		})
		// Zero-weight edge from the source to the start of the job.
		dag.Add(&weighted.Edge{
			V: sourceIndex,
			W: jobStartIndex,
		})
		// Zero-weight edge from the finish of the job to the sink.
		dag.Add(&weighted.Edge{
			V: jobFinishIndex,
			W: sinkIndex,
		})
		// Zero-weight edge for precedence constraints: from the end vertex corresponding to v
		// to the beginning vertex corresponding to w.
		for _, sucessor := range jobConstraints {
			dag.Add(&weighted.Edge{
				V: jobFinishIndex,
				W: sucessor,
			})
		}

		jobStartIndex++
		jobFinishIndex = jobStartIndex + jobsCount
	}
	if err := scanner.Err(); err != nil {
		return -1, nil, fmt.Errorf("failed to read a digraph: %w", err)
	}

	lp := lpt.NewAcyclic(dag, sourceIndex)
	startTimes := make([]float64, jobsCount)
	for i := 0; i < jobsCount; i++ {
		startTimes[i] = lp.DistTo(i)
	}
	finishTime := lp.DistTo(sinkIndex)
	return finishTime, startTimes, nil
}

// parseJobConstraint parses a string like 41.0 1 7 9 and
// returns a job duration (41.0) and a slice of job indexes which must complete before the given job.
func parseJobConstraint(s string) (float64, []int, error) {
	ss := strings.Fields(s)
	if len(ss) == 0 {
		return -1, nil, fmt.Errorf("empty job")
	}

	duration, err := strconv.ParseFloat(ss[0], 64)
	if err != nil {
		return -1, nil, fmt.Errorf("invalid duration format: %w", err)
	}

	// The job doesn't have constraints.
	if len(ss) == 1 {
		return duration, nil, nil
	}

	jobs := make([]int, 0, len(ss)-1)
	for i := 1; i < len(ss); i++ {
		id, err := strconv.Atoi(ss[i])
		if err != nil {
			return -1, nil, fmt.Errorf("invalid job index format: %w", err)
		}
		jobs = append(jobs, id)
	}

	return duration, jobs, nil
}
