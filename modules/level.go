package modules

import (
	"github.com/Estuardo2015/rogue_wizard/commons"
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

func NewLevel(w, h, tw int, p *Player) *Level {
	l := &Level{
		Width:     w, // tiles wide
		Height:    h, // tiles high
		TileWidth: tw,

		TileGrid: make([][]*Tile, w),

		Player: p,
	}

	// Initialize inner slices
	for i := range l.TileGrid {
		l.TileGrid[i] = make([]*Tile, l.Height)
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
