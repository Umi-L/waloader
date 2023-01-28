package waloader

import (
	"testing"
)

func TestLoad(t *testing.T) {

	sprites := LoadAtlas("assets/atlases/", "atlas.xml")

	_, ok := sprites["Test-1"]

	if !ok {
		t.Error("Test-1 not loaded!")
	}

	marioSprite, ok := sprites["MarioSpriteSheet"]

	if !ok {
		t.Error("MarioSpriteSheet not loaded!")
	}

	marioSheet := LoadSheet(marioSprite, 16, 32)

	LoadAnimation(&marioSheet, 0, 5, 1)
}
