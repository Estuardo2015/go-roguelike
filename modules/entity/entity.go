package entity

import (
	"github.com/Estuardo2015/rogue_wizard/modules/grid"
	"github.com/Estuardo2015/rogue_wizard/modules/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	GetPosition() (x, y int)
	Move(x, y int, g *grid.Grid)
	Image() *ebiten.Image
}

func MoveEntities(entities []Entity, grid *grid.Grid) {
	for _, e := range entities {
		x, y := e.GetPosition()
		switch pickDirection() {
		case 0: // Up
			e.Move(x, y-1, grid)
		case 1: // Down
			e.Move(x, y+1, grid)
		case 2: // Left
			e.Move(x-1, y, grid)
		case 3: // Right
			e.Move(x+1, y, grid)
		}
	}
}

func pickDirection() int {
	return utils.RandInt(4)
}
