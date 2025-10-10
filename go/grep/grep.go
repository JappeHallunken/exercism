package grep

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Search(pattern string, flags, files []string) []string {
	fl := make(map[string]bool, len(flags)) 
	for _, f := range flags {
		fl[f] = true
	}

	showFilename := len(files) > 1 
	var results []string

	for _, fn := range files {
		file, err := os.Open(fn)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fehler bei %s: %v\n", fn, err)
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lineNum := 0
		sPattern := pattern 
		if fl["-i"] { 
			sPattern = strings.ToLower(pattern) 
		}

		for scanner.Scan() {
			lineNum++
			line := scanner.Text()
			sLine := line
			if fl["-i"] {
				sLine = strings.ToLower(line)
			}

			match := fl["-x"] && sLine == sPattern || !fl["-x"] && strings.Contains(sLine, sPattern)
			if fl["-v"] {
				match = !match
			}
			if !match {
				continue
			}

			if fl["-l"] {
				results = append(results, fn)
				break
			}

			var b strings.Builder
			if showFilename {
				b.WriteString(fn)
				b.WriteByte(':')
			}

			if fl["-n"] {
				b.WriteString(strconv.Itoa(lineNum))
				b.WriteByte(':')
			}
			b.WriteString(line)
			results = append(results, b.String())
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Fehler beim Lesen von %s: %v\n", fn, err)
		}
	}

	return results
}
