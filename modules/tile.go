package modules

import "github.com/hajimehoshi/ebiten/v2"

type Tile struct {
	X       int
	Y       int
	Image   *ebiten.Image
	Blocked bool
}

func NewTile(x, y int, i *ebiten.Image, b bool) *Tile {
	return &Tile{
		X:       x,
		Y:       y,
		Image:   i,
		Blocked: b,
	}
}
