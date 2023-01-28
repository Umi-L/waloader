package waloader

import (
	"encoding/xml"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type atlas struct {
	XMLName  xml.Name  `xml:"atlas"`
	Textures []texture `xml:"tex"`
}

type texture struct {
	Sprites []Sprite `xml:"img"`
	Path    string   `xml:"n,attr"`
}

func LoadAtlas(pathToAtlasFolder, atlasName string) map[string]Sprite {
	dat, err := os.ReadFile(pathToAtlasFolder + atlasName)
	if err != nil {
		panic(err)
	}

	var at atlas
	if err := xml.Unmarshal(dat, &at); err != nil {
		panic(err)
	}

	var sprites map[string]Sprite

	sprites = make(map[string]Sprite)

	for t := range at.Textures {
		tex := at.Textures[t]

		img, _, err := ebitenutil.NewImageFromFile(pathToAtlasFolder + tex.Path + ".png")
		if err != nil {
			panic(err)
		}

		for i := range tex.Sprites {
			sprite := tex.Sprites[i]

			sprite.Rect = image.Rect(sprite.X, sprite.Y, sprite.X+sprite.Width, sprite.Y+sprite.Height)
			sprite.Image = img.SubImage(sprite.Rect).(*ebiten.Image)

			sprites[sprite.Name] = sprite
		}
	}

	return sprites
}

func LoadSheet(sprite Sprite, cellWidth int, cellHeight int) Sheet {
	width, height := sprite.Image.Size()

	return Sheet{
		Sprite:      sprite,
		CellWidth:   cellWidth,
		CellHeight:  cellHeight,
		Rows:        height / cellHeight,
		CellsPerRow: width / cellWidth,
	}
}

func LoadAnimation(parent *Sheet, index int, frames int, speed float32) Animation {
	var sprites []image.Rectangle

	for i := 0; i < frames; i++ {
		sprites = append(sprites, image.Rectangle{
			Min: image.Point{
				X: parent.CellWidth * i,
				Y: parent.CellHeight * index,
			},
			Max: image.Point{
				X: (parent.CellWidth * i) + parent.CellWidth,
				Y: (parent.CellHeight * index) + parent.CellHeight,
			},
		})
	}

	return Animation{
		parent:       parent,
		index:        index,
		frames:       frames,
		sprites:      sprites,
		speed:        speed,
		timer:        0.0,
		currentFrame: 0,
	}
}
