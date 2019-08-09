package strsort

import "fmt"

func ExampleLSD() {
	plates := []string{
		"4PGC938",
		"2IYE230",
		"3CI0720",
		"1ICK750",
		"1OHV845",
		"4JZY524",
		"1ICK750",
		"3CI0720",
		"1OHV845",
		"1OHV845",
		"2RLA629",
		"2RLA629",
		"3ATW723",
	}

	LSD(plates, 256)
	for i := range plates {
		fmt.Println(plates[i])
	}
	// Output:
	// 1ICK750
	// 1ICK750
	// 1OHV845
	// 1OHV845
	// 1OHV845
	// 2IYE230
	// 2RLA629
	// 2RLA629
	// 3ATW723
	// 3CI0720
	// 3CI0720
	// 4JZY524
	// 4PGC938
}
