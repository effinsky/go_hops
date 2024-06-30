package main

import (
	"fmt"

	"gohops/hops"
)

func main() {
	inp := `
    2 
    5 5
    0 0 4 4
    1
    2 2 2 2
    3 3
    0 0 2 2
    1
    1 1 1 1
    `
	cases := hops.MustParseInput(inp)
	results := hops.MinHops(cases...)
	fmt.Printf("results: %v\n", results)
}
