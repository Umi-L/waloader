package waloader

import (
	"encoding/xml"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type atlas struct {
	XMLName xml.Name `xml:"tex"`
	Sprites []Sprite `xml:"img"`
}

func LoadAtlas(imagePath, defPath string) []Sprite {
	dat, err := os.ReadFile(defPath)
	if err != nil {
		panic(err)
	}

	var atlas atlas
	if err := xml.Unmarshal(dat, &atlas); err != nil {
		panic(err)
	}

	img, _, err := ebitenutil.NewImageFromFile(imagePath, ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}

	for i := range atlas.Sprites {
		sprite := &atlas.Sprites[i]

		sprite.Rect = image.Rect(sprite.X, sprite.Y, sprite.X+sprite.Width, sprite.Y+sprite.Height)
		sprite.image = img.SubImage(sprite.Rect).(*ebiten.Image)
	}

	return atlas.Sprites
}
