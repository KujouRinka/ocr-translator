package scanner

import "image"

type Scanner interface {
	Scan() (image.Image, error)
	SetX(int) error
	SetY(int) error
	SetWidth(int) error
	SetHeight(int) error
}
