package modules

import (
	_ "image/png"
)

type Level struct {
	Width      int
	Height     int
	TileWidth  int
	TileHeight int

	TileMap   [][]*Tile
	EntityMap [][]*Entity

	Player *Entity
	Mobs   []*Entity
}

func NewLevel() *Level {
	s := &Level{
		Width:      40, // tiles wide
		Height:     25, // tiles high
		TileWidth:  32,
		TileHeight: 32,
	}
	PopulateTiles(s)

	return s
}

func PopulateTiles(s *Level) {
	s.TileMap = make([][]*Tile, s.Width)
	s.EntityMap = make([][]*Entity, s.Width)
	for i := range s.TileMap {
		s.TileMap[i] = make([]*Tile, s.Height)
		s.EntityMap[i] = make([]*Entity, s.Height)
	}

	for x := 0; x < s.Width; x++ {
		for y := 0; y < s.Height; y++ {
			// If on edge of map
			if x == 0 || x == s.Width-1 || y == 0 || y == s.Height-1 {
				tile := &Tile{
					X:       x,
					Y:       y,
					Image:   WallImg,
					Blocked: true,
				}
				s.TileMap[x][y] = tile
			}

			// Place Player
			if x == 20 && y == 20 {
				player := &Entity{
					Name:  "Player",
					Image: PlayerImg,
					X:     x,
					Y:     y,
				}
				s.EntityMap[x][y] = player
				s.Player = player
			}
		}
	}
}
