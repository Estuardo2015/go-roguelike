package modules

import (
	"github.com/Estuardo2015/rogue_wizard/modules/commons"
	_ "image/png"
)

type Level struct {
	Width     int
	Height    int
	TileWidth int

	TileGrid [][]*Tile

	Player   *Player
	Entities []Entity

	Camera *Camera
}

func NewLevel(w, h, tw int, tg [][]*Tile, p *Player) *Level {
	l := &Level{
		Width:     w, // tiles wide
		Height:    h, // tiles high
		TileWidth: tw,

		TileGrid: tg,

		Player: p,
	}

	// Initialize camera
	l.Camera = NewCamera(l.Player, commons.ScreenWidthTiles, commons.ScreenHeightTiles, l.Width, l.Height)

	return l
}

func BuildDemoLevel(l *Level) {
	for x := 0; x < l.Width; x++ {
		for y := 0; y < l.Height; y++ {
			// If on edge of map
			if x == 0 || x == l.Width-1 || y == 0 || y == l.Height-1 {
				tile := &Tile{
					X:       x,
					Y:       y,
					Image:   WallImg,
					Blocked: true,
				}
				l.TileGrid[x][y] = tile
			}
		}
	}
}
