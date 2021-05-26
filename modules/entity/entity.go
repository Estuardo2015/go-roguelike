package entity

import (
	"github.com/Estuardo2015/rogue_wizard/modules/grid"
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	GetPosition() (x, y int)
	Move(x, y int, g *grid.Grid)
	Image() *ebiten.Image
}
