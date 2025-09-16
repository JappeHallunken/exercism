package isogram

import "strings"

func IsIsogram(word string) bool {

	charMap := make(map[rune]bool) // default value is false
	word = strings.ToLower(word)

	for _, char := range word {
		if char >= 'a' && char <= 'z' { // Consider only alphabetic characters
			if charMap[char] {
				return false // Character already seen
			}
			charMap[char] = true // Mark character as seen
		}
	}
	return true // No duplicates found
}
