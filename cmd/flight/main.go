// Program flight takes a source vertex (JFK airport),
// a query (LAS) as arguments, and prints a shortest path (degrees of separation)
// from the source to the query vertex using breadth-first search.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/marselester/alg/graph"
)

func main() {
	origin := flag.String("from", "", "airport to search a shortest path from (origin)")
	destin := flag.String("to", "", "airport to search a shortest path to (destination)")
	sep := flag.String("sep", " ", "connections (edges) separator in a symbol graph")
	flag.Parse()
	if len(*origin) == 0 {
		log.Fatal("flight: search origin must not be blank")
	}
	if len(*destin) == 0 {
		log.Fatal("flight: search destination must not be blank")
	}

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("flight: read failed: %v", err)
	}
	r := bytes.NewReader(b)

	sg, err := graph.NewSymbolGraph(r, *sep)
	if err != nil {
		log.Fatalf("flight: graph creation failed: %v", err)
	}

	src, ok := sg.Index(*origin)
	if !ok {
		log.Fatalf("flight: origin not found")
	}
	dst, ok := sg.Index(*destin)
	if !ok {
		log.Fatalf("flight: destination not found")
	}

	bfs := graph.NewBreadthFirstPath(sg.Graph(), src)
	for _, v := range bfs.To(dst) {
		fmt.Println(sg.Name(v))
	}
}
