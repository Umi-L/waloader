package waloader

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

type Sprite struct {
	Name   string `xml:"n,attr"`
	X      int    `xml:"x,attr"`
	Y      int    `xml:"y,attr"`
	Width  int    `xml:"w,attr"`
	Height int    `xml:"h,attr"`
	image  *ebiten.Image
	Rect   image.Rectangle
}
