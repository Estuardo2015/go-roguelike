package modules

import (
	"github.com/Estuardo2015/rogue_wizard/modules/commons"
	_ "image/png"
)

type Level struct {
	Width     int // tiles wide
	Height    int // tiles high
	TileWidth int // tile pixel width

	TileGrid [][]*Tile

	Player   *Player
	Entities []Entity

	Camera *Camera
}

func NewLevel(w, h, tw int, tg [][]*Tile, p *Player) *Level {
	l := &Level{
		Width:     w,
		Height:    h,
		TileWidth: tw,

		TileGrid: tg,

		Player: p,
	}

	// Initialize camera
	l.Camera = NewCamera(l.Player, commons.ScreenWidthTiles, commons.ScreenHeightTiles, l.Width, l.Height)

	return l
}
