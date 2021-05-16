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
