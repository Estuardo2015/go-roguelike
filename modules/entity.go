package modules

import "github.com/hajimehoshi/ebiten/v2"

type Entity interface {
	GetPosition() (x, y int)
	Move(x, y int, tg [][]*Tile)
}

type Mob struct {
	Name    string
	X       int
	Y       int
	Image   *ebiten.Image
	Blocked bool
}

func (m Mob) GetPosition() (x, y int) {
	return m.X, m.Y
}
