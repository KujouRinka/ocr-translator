package ocr

import (
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)

func Screenshot() {
	bounds := screenshot.GetDisplayBounds(0)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("screenshot.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}
