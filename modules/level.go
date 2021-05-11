package modules

import (
	_ "image/png"
)

type Level struct {
	Width      int
	Height     int
	TileWidth  int
	TileHeight int

	TileGrid [][]*Tile

	Player   *Entity
	Entities []*Entity
}

func NewLevel(w, h, tw, th, pX, pY int) *Level {
	l := &Level{
		Width:      w, // tiles wide
		Height:     h, // tiles high
		TileWidth:  tw,
		TileHeight: th,

		TileGrid: make([][]*Tile, w),

		Player: &Entity{
			Name:  "Player",
			Image: PlayerImg,
			X:     pX,
			Y:     pY,
		},
	}

	// Initialize inner slices
	for i := range l.TileGrid {
		l.TileGrid[i] = make([]*Tile, l.Height)
	}

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
