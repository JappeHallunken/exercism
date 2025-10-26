package markdown

import (
	"fmt"
	"regexp"
	"strings"
)

func parseInline(s string) string {
	s = regexp.MustCompile(`__(.*?)__`).ReplaceAllString(s, "<strong>$1</strong>")
	s = regexp.MustCompile(`_(.*?)_`).ReplaceAllString(s, "<em>$1</em>")
	return s
}

func parseHeader(line string) (string, bool) {
	n := 0
	for n < len(line) && line[n] == '#' {
		n++
	}
	if n == 0 || n > 6 {
		return "", false
	}
	content := strings.TrimSpace(line[n:])
	return fmt.Sprintf("<h%d>%s</h%d>", n, parseInline(content), n), true
}

func parseList(lines []string) string {
	var html strings.Builder
	html.WriteString("<ul>")
	for _, line := range lines {
		content := strings.TrimPrefix(line, "* ")
		html.WriteString("<li>" + parseInline(content) + "</li>")
	}
	html.WriteString("</ul>")
	return html.String()
}

func Render(markdown string) string {
	lines := strings.Split(strings.TrimSpace(markdown), "\n")
	var html strings.Builder
	var list []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if h, ok := parseHeader(line); ok {
			if len(list) > 0 {
				html.WriteString(parseList(list))
				list = nil
			}
			html.WriteString(h)
			continue
		}
		if strings.HasPrefix(line, "* ") {
			list = append(list, line)
			continue
		}
		if len(list) > 0 {
			html.WriteString(parseList(list))
			list = nil
		}
		html.WriteString("<p>" + parseInline(line) + "</p>")
	}
	if len(list) > 0 {
		html.WriteString(parseList(list))
	}
	return html.String()
}
