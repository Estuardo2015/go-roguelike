package modules

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Name string
	X    int
	Y    int

	Image *ebiten.Image
}

func NewPlayer(n string, x, y int, i *ebiten.Image) *Player {
	return &Player{
		Name:  n,
		X:     x,
		Y:     y,
		Image: i,
	}
}

func (pl *Player) GetPosition() (x, y int) {
	return pl.X, pl.Y
}

func (pl *Player) Move(x int, y int, tg [][]*Tile) {
	if tg[x][y] == nil || !tg[x][y].Blocked {
		pl.X = x
		pl.Y = y
	}
}