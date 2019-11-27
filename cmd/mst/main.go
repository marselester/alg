// Program mst reads an edge-weighted graph from stdin, computes the MST
// (minimum spanning tree) of that graph, prints the MST edges,
// and prints the total weight of the MST.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/marselester/alg/graph/mst"
)

func main() {
	vertices := flag.Int("v", 0, "number of vertices in the graph")
	alg := flag.String("alg", "prim", "algorithm to compute MST (prim, kruskal)")
	flag.Parse()

	g := mst.NewAdjacencyList(*vertices)
	// Lines from stdin such as 4 5 0.35 represent 4-5 edge with 0.35 weight.
	var v, w int
	var weight float64
	var err error

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		vv := strings.Fields(scanner.Text())
		if len(vv) != 3 {
			log.Printf("mst: expected two vertices and weight, got %d", len(vv))
			continue
		}

		if v, err = strconv.Atoi(vv[0]); err != nil {
			log.Fatalf("mst: invalid vertex format: %v", err)
		}
		if v >= *vertices || v < 0 {
			log.Fatalf("mst: vertex %d is out of range %d", v, *vertices)
		}

		if w, err = strconv.Atoi(vv[1]); err != nil {
			log.Fatalf("mst: invalid vertex format: %v", err)
		}
		if w >= *vertices || w < 0 {
			log.Fatalf("mst: vertex %d is out of range %d", w, *vertices)
		}

		if weight, err = strconv.ParseFloat(vv[2], 32); err != nil {
			log.Fatalf("mst: invalid weight format: %v", err)
		}

		g.Add(&mst.Edge{
			V:      v,
			W:      w,
			Weight: float32(weight),
		})
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("mst: failed to read a graph: %v", err)
	}

	var tree interface {
		Edges() []*mst.Edge
		Weight() float32
	}
	switch *alg {
	case "kruskal":
		tree = mst.NewKruskal(g)
	default:
		tree = mst.NewLazyPrim(g)
	}

	for _, e := range tree.Edges() {
		fmt.Println(e)
	}
	fmt.Printf("%.5f\n", tree.Weight())
}
