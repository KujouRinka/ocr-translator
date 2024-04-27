package translator

import (
	"testing"

	"golang.org/x/text/language"
)

const (
	APIKey = "FILL YOUR API HERE"
)

func TestGoogleTranslator(t *testing.T) {
	gt, err := NewGoogleTranslator(language.SimplifiedChinese, language.Japanese, APIKey)
	if err != nil {
		t.Fatal(err)
	}

	text := "こんにちは"
	translated, err := gt.Translate(text)
	if err != nil {
		t.Fatal(err)
	}

	if translated == "" {
		t.Fatal("translated text is empty")
	}

	t.Logf("translated text: %s", translated)
}
