package ocr

import (
	"image/png"
	"os"
	"testing"
)

func TestGosseractOCR(t *testing.T) {
	ocr, err := NewGosseractOcr("jpn")
	if err != nil {
		t.Fatal(err)
	}
	defer ocr.Close()

	// Test image is "testdata/test.png"
	imgFile, err := os.Open("testdata/test.png")
	if err != nil {
		t.Fatal(err)
	}
	defer imgFile.Close()

	img, err := png.Decode(imgFile)
	if err != nil {
		t.Fatal(err)
	}

	text, err := ocr.ImgToText(img)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(text)
}
