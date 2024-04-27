package ocr

import "strings"

func Format(text string) string {
	return strings.ReplaceAll(text, " ", "")
}
