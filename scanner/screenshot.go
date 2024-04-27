package scanner

import (
	"image"

	"github.com/kbinani/screenshot"
)

type DefaultScanner struct {
	x      int
	y      int
	width  int
	height int
}

func NewDefaultScanner(x, y, width, height int) *DefaultScanner {
	return &DefaultScanner{x: x, y: y, width: width, height: height}
}

func (s *DefaultScanner) Scan() (image.Image, error) {
	var bounds image.Rectangle
	if s.x == -1 || s.y == -1 || s.width == -1 || s.height == -1 {
		bounds = screenshot.GetDisplayBounds(0)
	} else {
		bounds = image.Rect(s.x, s.y, s.x+s.width, s.y+s.height)
	}
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func (s *DefaultScanner) SetX(x int) error {
	s.x = x
	return nil
}

func (s *DefaultScanner) SetY(y int) error {
	s.y = y
	return nil
}

func (s *DefaultScanner) SetWidth(width int) error {
	s.width = width
	return nil
}

func (s *DefaultScanner) SetHeight(height int) error {
	s.height = height
	return nil
}
