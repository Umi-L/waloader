package waloader

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	parent       *Sheet
	index        int
	frames       int
	sprites      []image.Rectangle
	speed        float32
	timer        float32
	currentFrame int
}

func (anim *Animation) Draw(screen *ebiten.Image, x float64, y float64, rotation float64) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(x, y)
	op.GeoM.Rotate(rotation)

	//fmt.Printf("Frame Rects: %+v \n", anim.sprites)
	//fmt.Printf("Image: %+v \n", anim.parent.Sprite.Image)

	offset := anim.sprites[anim.currentFrame]
	offset.Min.X += anim.parent.Sprite.Image.Bounds().Min.X
	offset.Min.Y += anim.parent.Sprite.Image.Bounds().Min.Y

	offset.Max.X += anim.parent.Sprite.Image.Bounds().Min.X
	offset.Max.Y += anim.parent.Sprite.Image.Bounds().Min.Y

	screen.DrawImage(anim.parent.Sprite.Image.SubImage(offset).(*ebiten.Image), op)
}

func (anim *Animation) UpdateTimer(dt float32) {
	anim.timer += dt

	if anim.timer > anim.speed {
		anim.currentFrame += 1
		anim.timer = 0.0
	}

	if anim.currentFrame == anim.frames {
		anim.currentFrame = 0
	}
}
