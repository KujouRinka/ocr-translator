package ocr

import (
	"regexp"
	"strings"
)

var (
	re = regexp.MustCompile(`\s{5,}`)
)

func Format(text string) string {
	lines := strings.Split(text, "\n")
	newlines := make([]string, 0, len(lines)/2)
	// remove empty lines and line without valid characters
	for _, line := range lines {
		line = re.ReplaceAllString(line, "    ")
		if hasInvalidChars(line) {
			continue
		}
		newlines = append(newlines, line)
	}

	return strings.Join(newlines, "\n")
}

func hasInvalidChars(s string) bool {
	var invalidChar int
	for _, r := range []rune(s) {
		if r != ' ' && (r <= 0x2F || r >= 0x3A && r <= 0x40 || r >= 0x5B && r <= 0x60 || r >= 0x7B && r <= 0x7E) {
			invalidChar++
		}
	}

	return invalidChar*2 > len(s)
}
