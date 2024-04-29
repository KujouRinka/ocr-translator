package scanner

import "image"

type Scanner interface {
	Scan() (image.Image, error)
	GetX() int
	SetX(int) error
	GetY() int
	SetY(int) error
	GetWidth() int
	SetWidth(int) error
	GetHeight() int
	SetHeight(int) error
}
