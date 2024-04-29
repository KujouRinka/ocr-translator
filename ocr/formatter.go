package ocr

import (
	"strings"
)

func Format(text string) string {
	lines := strings.Split(text, "\n")
	newlines := make([]string, 0, len(lines)/2)
	// remove empty lines and line without valid characters
	for _, line := range lines {
		if !hasValidChars(line) {
			continue
		}
		newlines = append(newlines, line)
	}

	return strings.Join(lines, "\n")
}

func hasValidChars(s string) bool {
	var invalidChar int
	var spaceCnt int
	for _, r := range []rune(s) {
		if r == ' ' {
			spaceCnt++
		} else if r <= 0x2F || r >= 0x3A && r <= 0x40 || r >= 0x5B && r <= 0x60 || r >= 0x7B && r <= 0x7E {
			invalidChar++
		}
	}

	return (spaceCnt+invalidChar)*5 > len(s)*4
}
