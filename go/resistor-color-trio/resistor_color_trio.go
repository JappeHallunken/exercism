package resistorcolortrio

import (
	"math"
	"strconv"
)

var col = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

const (
	Kilo = 1_000
	Mega = 1_000_000
	Giga = 1_000_000_000
)

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(s []string) string {
	fi := col[s[0]]
	si := col[s[1]]
	e := col[s[2]]

	if fi == -1 || si == -1 || e == -1 || len(s) < 3 {
		return ""
	}
	n := (fi*10 + si) * int(math.Pow10(e))

	switch {
	case n >= Giga:
		n /= Giga
		return strconv.Itoa(n) + " gigaohms"
	case n >= Mega:
		n /= Mega
		return strconv.Itoa(n) + " megaohms"
	case n >= Kilo:
		n /= Kilo
		return strconv.Itoa(n) + " kiloohms"
	default:
		return strconv.Itoa(n) + " ohms"
	}
}
