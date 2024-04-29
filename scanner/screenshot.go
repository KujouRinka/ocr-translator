package scanner

import (
	"image"
	"sync"

	"github.com/kbinani/screenshot"
)

var _ Scanner = (*DefaultScanner)(nil)

type DefaultScanner struct {
	x      int
	y      int
	width  int
	height int

	mu sync.RWMutex
}

func NewDefaultScanner(x, y, width, height int) *DefaultScanner {
	return &DefaultScanner{x: x, y: y, width: width, height: height}
}

func (s *DefaultScanner) Scan() (image.Image, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

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

func (s *DefaultScanner) GetX() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.x
}

func (s *DefaultScanner) SetX(x int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.x = x
	return nil
}

func (s *DefaultScanner) GetY() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.y
}

func (s *DefaultScanner) SetY(y int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.y = y
	return nil
}

func (s *DefaultScanner) GetWidth() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.width
}

func (s *DefaultScanner) SetWidth(width int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if width < 0 {
		width = 0
	}
	s.width = width
	return nil
}

func (s *DefaultScanner) GetHeight() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.height
}

func (s *DefaultScanner) SetHeight(height int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if height < 0 {
		height = 0
	}
	s.height = height
	return nil
}
