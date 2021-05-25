package modules

import (
	"github.com/Estuardo2015/rogue_wizard/modules/components"
	"github.com/Estuardo2015/rogue_wizard/modules/grid"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Name string
	components.PositionComponent
	components.HealthComponent
	components.MagicComponent
	components.InventoryComponent

	image *ebiten.Image
}

func NewPlayer(n string, x, y int, i *ebiten.Image) *Player {
	p := &Player{
		Name:  n,
		image: i,
	}
	p.X = x
	p.Y = y

	return p
}

func (pl *Player) GetPosition() (x, y int) {
	return pl.X, pl.Y
}

func (pl *Player) Move(x int, y int, g *grid.Grid) {
	if g.TileAt(x, y) == nil || !g.TileAt(x, y).Blocked {
		pl.X = x
		pl.Y = y
	}
	//log.Debug().Msgf("Move - PlayerX: %d - PlayerY: %d", pl.X, pl.Y)
}

func (pl *Player) Image() *ebiten.Image {
	return pl.image
}
