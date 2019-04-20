// Program graph reads a graph from the input stream and then prints it
// to show the order in which vertices appear in adjacency lists.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/marselester/alg/graph"
)

func main() {
	vertices := flag.Int("v", 0, "number of vertices in the graph")
	flag.Parse()
	g := graph.NewAdjacencyList(*vertices)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		vv := strings.Fields(scanner.Text())
		if len(vv) != 2 {
			log.Printf("graph: expected two vertices, got %d", len(vv))
			continue
		}

		v, err := strconv.Atoi(vv[0])
		if err != nil {
			log.Fatalf("graph: invalid vertex format: %v", err)
		}
		if v >= *vertices || v < 0 {
			log.Fatalf("graph: vertex %d is out of range %d", v, *vertices)
		}
		w, err := strconv.Atoi(vv[1])
		if err != nil {
			log.Fatalf("graph: invalid vertex format: %v", err)
		}
		if w >= *vertices || w < 0 {
			log.Fatalf("graph: vertex %d is out of range %d", w, *vertices)
		}

		g.Add(v, w)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("graph: failed to read a graph: %v", err)
	}

	fmt.Print(g)
}
