package anagram

import "strings"

func Detect(subject string, candidates []string) []string {

	subject = strings.ToLower(subject)
	result := []string{}

	for _, candidate := range candidates {
		if len(candidate) != len(subject) {
			continue
		}
		candidateLower := strings.ToLower(candidate)
		//skip if the candidate is the same as the subject
		if candidateLower == subject {
			continue
		}

		subjectRunes := []rune(subject)
		candidateRunes := []rune(candidateLower)

		if isAnagram(subjectRunes, candidateRunes) {
			result = append(result, candidate)
		}
	}

	return result
}

func isAnagram(subject, candidate []rune) bool {

	if len(subject) != len(candidate) {
		return false
	}

	for _, s := range subject {
		for j, c := range candidate {
			if s == c {
				candidate = append(candidate[:j], candidate[j+1:]...)
				break
			}
		}
	}

	return len(candidate) == 0
}
