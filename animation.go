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

	screen.DrawImage(anim.parent.Sprite.Image.SubImage(anim.sprites[anim.currentFrame]).(*ebiten.Image), op)
}

func (anim *Animation) updateTimer(dt float32) {
	anim.timer += dt

	if anim.timer > anim.speed {
		anim.currentFrame += 1
		anim.timer = 0.0
	}

	if anim.currentFrame == anim.frames {
		anim.currentFrame = 0
	}
}
