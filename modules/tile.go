package modules

import "github.com/hajimehoshi/ebiten/v2"

type Tile struct {
	X       int
	Y       int
	Image   *ebiten.Image
	Blocked bool
}

func GetTile(l *Level, x, y int) *Tile {
	if x < 0 || x >= l.Width || y < 0 || y >= l.Height {
		return nil
	}
	return l.TileGrid[x][y]
}

func NewTile(x, y int, i *ebiten.Image, b bool) *Tile {
	return &Tile{
		X:       x,
		Y:       y,
		Image:   i,
		Blocked: b,
	}
}

func NewTileGrid() [][]*Tile {
	level := make([][]*Tile, 0)
	for i, _ := range level {
		level[i] = make([]*Tile, 0)
	}

	return level
}
