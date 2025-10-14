package foodchain

import (
	"fmt"
	"strings"
)

type animal struct {
	name      string
	comment   string
	catchText string
}

var animals = []animal{
	{"fly", "", ""},
	{"spider", "It wriggled and jiggled and tickled inside her.", "that wriggled and jiggled and tickled inside her"},
	{"bird", "How absurd to swallow a bird!", ""},
	{"cat", "Imagine that, to swallow a cat!", ""},
	{"dog", "What a hog, to swallow a dog!", ""},
	{"goat", "Just opened her throat and swallowed a goat!", ""},
	{"cow", "I don't know how she swallowed a cow!", ""},
	{"horse", "She's dead, of course!", ""},
}

func Verse(v int) string {
	var sb strings.Builder
	a := animals[v-1]

	sb.WriteString(fmt.Sprintf("I know an old lady who swallowed a %s.\n", a.name))

	if a.name == "horse" {
		if a.comment != "" {
			sb.WriteString(a.comment) 
		}
		return sb.String()
	}

	if a.comment != "" {
		sb.WriteString(a.comment + "\n")
	}

	for i := v - 1; i > 0; i-- {
		current := animals[i]
		caught := animals[i-1]

		sb.WriteString(fmt.Sprintf("She swallowed the %s to catch the %s", current.name, caught.name))
		if caught.catchText != "" {
			sb.WriteString(" " + caught.catchText)
		}
		sb.WriteString(".\n")
	}

	sb.WriteString("I don't know why she swallowed the fly. Perhaps she'll die.")

	return sb.String()
}

func Verses(start, end int) string {
	var verses []string
	for i := start; i <= end; i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n\n")
}

func Song() string {
	return Verses(1, len(animals))
}

func main() {
	fmt.Println(Song())
}
