// Package proverb generates a proverb from a list of words.
package proverb

import "fmt"

// Proverb generates a proverb from the input slice of strings.
func Proverb(rhyme []string) []string {

	if len(rhyme) == 0 {
		return []string{}
	}

	var result []string

	for i:=0; i < len(rhyme)-1; i++ {
		result = append(result, produceLine(rhyme[i], rhyme[i+1]))
	}
	result = append(result, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return result
}

func produceLine(a, b string) string {
	return fmt.Sprintf("For want of a %s the %s was lost.", a, b)
}
