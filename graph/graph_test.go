package graph_test

import (
	"bytes"
	"fmt"
	"log"

	"github.com/marselester/alg/graph"
)

const routes = `
JFK MCO
ORD DEN
ORD HOU
DFW PHX
JFK ATL
ORD DFW
ORD PHX
ATL HOU
DEN PHX
PHX LAX
JFK ORD
DEN LAS
DFW HOU
ORD ATL
LAS LAX
ATL MCO
HOU MCO
LAS PHX
`

func ExampleSymbolGraph() {
	r := bytes.NewReader([]byte(routes))

	sg, err := graph.NewSymbolGraph(r, " ")
	if err != nil {
		log.Fatalf("flight: graph creation failed: %v", err)
	}

	v, ok := sg.Index("JFK")
	if !ok {
		log.Fatalf("flight: airport not found")
	}
	for _, v := range sg.Adjacent(v) {
		fmt.Println(sg.Name(v))
	}
	// Output:
	// ORD
	// ATL
	// MCO
}
