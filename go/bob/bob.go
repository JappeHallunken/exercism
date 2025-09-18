// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {

	noSpaces := strings.TrimSpace(strings.ReplaceAll(remark, " ", ""))

	switch {
	case strings.HasSuffix(noSpaces, "?") && strings.ToUpper(noSpaces) == noSpaces && strings.ToLower(noSpaces) != noSpaces:
		return "Calm down, I know what I'm doing!"
	case strings.ToUpper(noSpaces) == noSpaces && strings.ToLower(noSpaces) != noSpaces:
		return "Whoa, chill out!"
	case strings.HasSuffix(noSpaces, "?"):
		return "Sure."
	case noSpaces == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
