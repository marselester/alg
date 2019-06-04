// Program imdb looks up a movie or a performer by its name using a symbol graph.
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
	query := flag.String("query", "", "movie or performer to look for in standard input")
	sep := flag.String("sep", "/", "edges separator in a symbol graph")
	flag.Parse()
	if len(*query) == 0 {
		log.Fatal("imdb: search query must not be blank")
	}

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("imdb: read failed: %v", err)
	}
	r := bytes.NewReader(b)

	sg, err := graph.NewSymbolGraph(r, *sep)
	if err != nil {
		log.Fatalf("imdb: graph creation failed: %v", err)
	}

	v, ok := sg.Index(*query)
	if !ok {
		log.Fatalf("imdb: movie or performer not found")
	}
	for _, v := range sg.Adjacent(v) {
		fmt.Println(sg.Name(v))
	}
}
