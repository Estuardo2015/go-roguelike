package modules

import "github.com/hajimehoshi/ebiten/v2"

type Entity interface {
	GetPosition() (x, y int)
	Move(x, y int, tg [][]*Tile)
	Image() *ebiten.Image
}

type Mob struct {
	Name    string
	X       int
	Y       int
	Blocked bool

	image *ebiten.Image
}

func (m *Mob) GetPosition() (x, y int) {
	return m.X, m.Y
}

func (m *Mob) Move(x int, y int, tg [][]*Tile) {
	if tg[x][y] == nil || !tg[x][y].Blocked {
		m.X = x
		m.Y = y
	}
}

func (m *Mob) Image() *ebiten.Image {
	return m.image
}
