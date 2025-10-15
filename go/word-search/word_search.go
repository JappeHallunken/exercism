// Package wordsearch searches a square of letters for a given list of words in any direction
// returns coordinates for the starting and ending letter of a word. returns {-1, -1} for coordinates of not found words
package wordsearch

import (
	"fmt"
	"sync"
)

var directions = [8][2]int{
	{0, 1}, {1, 1}, {1, 0}, {1, -1}, // → ↘ ↓ ↙
	{0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, // ← ↖ ↑ ↗
}

// buildMatrix converts the slice of strings into a 2D rune matrix / square
func buildMatrix(puzzle []string) [][]rune {
	m := make([][]rune, len(puzzle))
	for i, s := range puzzle {
		m[i] = []rune(s)
	}
	return m
}

// wordExists checks if the word exists starting from (row, col).
// If found, returns true and the end coordinate of the word; otherwise returns false.
func wordExists(matrix [][]rune, word string, row, col int) (bool, [2]int) {
	for _, dir := range directions {
		dr, dc := dir[0], dir[1]
		found := true
		for i := range word {
			r := row + i*dr
			c := col + i*dc
			if r < 0 || r >= len(matrix) || c < 0 || c >= len(matrix[0]) || matrix[r][c] != rune(word[i]) {
				found = false
				break
			}
		}
		if found {
			return true, [2]int{col + (len(word)-1)*dc, row + (len(word)-1)*dr}
		}
	}
	return false, [2]int{}
}

// searchWord searches for the word in the matrix.
// It looks for the first rune and then checks all directions using wordExists.
// Returns start and end coordinates if found, otherwise [-1,-1].
func searchWord(matrix [][]rune, word string) [2][2]int {
	for row := range matrix {
		for col := range matrix[0] {
			if matrix[row][col] == rune(word[0]) {
				if ok, end := wordExists(matrix, word, row, col); ok {
					start := [2]int{col, row}
					return [2][2]int{start, end}
				}
			}
		}
	}
	return [2][2]int{{-1, -1}, {-1, -1}}
}

// Solve searches for all words in the puzzle concurrently.
// Returns a map of words to their start/end coordinates and an error if any word was not found.
// Example result: map["WORD"] = [[startCol, startRow], [endCol, endRow]]
func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	matrix := buildMatrix(puzzle)
	result := make(map[string][2][2]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, word := range words {
		wg.Add(1)
		go func(w string) {
			defer wg.Done()
			coords := searchWord(matrix, w)

			mu.Lock()
			result[w] = coords
			mu.Unlock()
		}(word)
	}

	wg.Wait()
	var err error = nil
	for w, c := range result {
		if c[0][0] == -1 {
			err = fmt.Errorf("word %q not found", w)
		}
	}
	return result, err
}
