package kindergarten

import (
	"fmt"
	"sort"
	"strings"
)

// Define the Garden type here.
type Garden map[string]string

var seeds = map[rune]string{
	'G': "grass",
	'C': "clover",
	'R': "radishes",
	'V': "violets",
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
	sortChild := append([]string{}, children...)
	sort.Strings(sortChild)
	rows := strings.Fields(diagram)

	if diagram[0] != '\n' {
		return nil, fmt.Errorf("invalid diagram format")
	}

	if len(rows) != 2 ||
		len(rows[0]) != len(rows[1]) ||
		2*len(sortChild) != len(rows[0]) {
		return nil, fmt.Errorf("invalid diagram")
	}

	g := make(Garden)

	for i, child := range sortChild {
		m := 2 * i // startindex per child for the cupcode; 0: [0:2], 1: [2:4], 2: [4,6] ...
		top, bottom := rows[0][m:m+2], rows[1][m:m+2]
		cupcode := top + bottom

		for _, r := range cupcode {
			if _, ok := seeds[r]; !ok {
				return nil, fmt.Errorf("invalid cup code")
			}
		}

		if _, exists := g[child]; exists {
			return nil, fmt.Errorf("duplicate name")
		}

		g[child] = cupcode
	}
	return &g, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := (*g)[child]
	if !ok {
		return nil, false
	}
	var result []string

	for _, r := range plants {
		if plant, found := seeds[r]; found {
			result = append(result, plant)
		}
	}
	return result, true
}
