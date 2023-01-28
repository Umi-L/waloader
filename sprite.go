package waloader

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Name   string `xml:"n,attr"`
	X      int    `xml:"x,attr"`
	Y      int    `xml:"y,attr"`
	Width  int    `xml:"w,attr"`
	Height int    `xml:"h,attr"`
	Image  *ebiten.Image
	Rect   image.Rectangle
}
