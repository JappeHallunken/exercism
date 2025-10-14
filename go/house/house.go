package house

import (
	"fmt"
	"strings"
)

type VersePart struct {
	Noun   string
	Clause string
}

var parts = []VersePart{
	{"the house that Jack built.", ""},
	{"the malt", "that lay in"},
	{"the rat", "that ate"},
	{"the cat", "that killed"},
	{"the dog", "that worried"},
	{"the cow with the crumpled horn", "that tossed"},
	{"the maiden all forlorn", "that milked"},
	{"the man all tattered and torn", "that kissed"},
	{"the priest all shaven and shorn", "that married"},
	{"the rooster that crowed in the morn", "that woke"},
	{"the farmer sowing his corn", "that kept"},
	{"the horse and the hound and the horn", "that belonged to"},
}

func Verse(v int) string {
	verse := "This is " + parts[v-1].Noun

	for j := v - 2; j >= 0; j-- {
		verse += "\n" + parts[j+1].Clause + " " + parts[j].Noun
	}
	return verse
}

func Song() string {
	var verses strings.Builder

	for i := 1; i <= len(parts); i++ {
		verses.WriteString(Verse(i))
		if i < len(parts) {
			verses.WriteString("\n\n")
		}
	}
	fmt.Println(verses.String())
	return verses.String()
}
