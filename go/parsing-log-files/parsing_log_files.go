package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*-=]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`^.*".*(?i)password.*"`)
	counter := 0
	for _, line := range lines {
		if re.MatchString(line) {
			counter++
		}
	}
	return counter
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	var result []string
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			username := matches[1]
			line = "[USR] " + username + " " + line
		}
		result = append(result, line)
	}
	return result
}
