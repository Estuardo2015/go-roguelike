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

func (l *Level) GetTile(x, y int) *Tile {
	if x < 0 || x >= l.Width || y < 0 || y >= l.Height {
		return nil
	}
	return l.TileGrid[x][y]
}

func (l *Level) AddEntity(e Entity) {
	l.Entities = append(l.Entities, e)
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
