package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	resultMap := FreqMap{}

	resultCh := make(chan FreqMap, len(texts))

	for _, t := range texts {
		go func(s string) { resultCh <- Frequency(s) }(t)
	}

	for range texts {
		for k, v := range <- resultCh {
			resultMap[k] += v
		}
	}
	return resultMap
}
