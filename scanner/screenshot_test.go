package scanner

import (
	"image/png"
	"os"
	"testing"
)

func TestScreenshot(t *testing.T) {
	s := NewDefaultScanner(383, 202, 695-383, 312-202)
	img, err := s.Scan()
	if err != nil {
		t.Fatalf("Failed to capture screenshot: %v", err)
	}
	if img == nil {
		t.Fatalf("Failed to capture screenshot: image is nil")
	}

	// img to png
	file, err := os.Create("testdata/screenshot.png")
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		t.Fatalf("Failed to encode image to png: %v", err)
	}
}
