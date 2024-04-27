package ocr

import "image"

type Engine interface {
	ImgToText(img image.Image) (string, error)
}
