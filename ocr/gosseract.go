package ocr

import (
	"bytes"
	"image"
	"image/png"

	"github.com/otiai10/gosseract/v2"
)

type GosseractOcr struct {
	client *gosseract.Client
}

func NewGosseractOcr(langs ...string) (*GosseractOcr, error) {
	client := gosseract.NewClient()

	err := client.SetLanguage(langs...)
	if err != nil {
		client.Close()
		return nil, err
	}

	return &GosseractOcr{
		client: client,
	}, nil
}

func (g *GosseractOcr) ImgToText(img image.Image) (string, error) {
	imgByte, err := imgToBytes(img)
	if err != nil {
		return "", nil
	}

	err = g.client.SetImageFromBytes(imgByte)
	if err != nil {
		return "", nil
	}

	text, err := g.client.Text()
	if err != nil {
		return "", nil
	}

	return text, nil
}

func imgToBytes(img image.Image) ([]byte, error) {
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (g *GosseractOcr) Close() error {
	return g.client.Close()
}
