package cipher

import "strings"

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.
type shift int
type vigenere string

const AL = 26

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance == 0 || distance <= -AL || distance >= AL {
		return nil
	}
	return shift(distance)
}

func (c shift) Encode(input string) string {
	var sb strings.Builder
	input = strings.ToLower(input)
	for _, r := range input {
		var shifted rune
		if r >= 'a' && r <= 'z' {
			offset := r - 'a'
			shifted = (offset + rune(c) + AL) % AL
			sb.WriteRune(shifted + 'a')
		}
	}
	return sb.String()
}

func (c shift) Decode(input string) string {
	return shift(-c).Encode(input)
}

func NewVigenere(key string) Cipher {
	if key == "" {
		return nil
	}
	hasNonA := false
	for _, r := range key {
		if r < 'a' || r > 'z' {
			return nil
		}
		if r != 'a' {
			hasNonA = true
		}
	}
	if !hasNonA {
		return nil
	}
	return vigenere(strings.ToLower(key))
}

func (v vigenere) apply(input string, shiftFunc func(int) shift) string {
	var sb strings.Builder
	keyIndex := 0
	keyLength := len(v)
	for _, r := range cleanString(input) {
		s := shiftFunc(getShift(rune(v[keyIndex%keyLength])))
		sb.WriteString(s.Encode(string(r)))
		keyIndex++
	}
	return sb.String()
}

func (v vigenere) Encode(input string) string {
	return v.apply(input, func(s int) shift { return shift(s) })
}
func (v vigenere) Decode(input string) string {
	return v.apply(input, func(s int) shift { return shift(-s) })
}

func getShift(r rune) int {
	return int(r - 'a')
}

// cleanString makes the string lowercase and removes everything except 'a' to 'z'
func cleanString(s string) string {
	var clean strings.Builder
	s = strings.ToLower(s)
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			clean.WriteRune(r)
		}
	}
	return clean.String()
}
