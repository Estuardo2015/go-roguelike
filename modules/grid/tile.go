package grid

import "github.com/hajimehoshi/ebiten/v2"

type Tile struct {
	Image   *ebiten.Image
	Blocked bool
}

func NewTile(i *ebiten.Image, b bool) *Tile {
	return &Tile{
		Image:   i,
		Blocked: b,
	}
}
