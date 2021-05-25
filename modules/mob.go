package modules

import (
	"github.com/Estuardo2015/rogue_wizard/modules/components"
	"github.com/Estuardo2015/rogue_wizard/modules/grid"
	"github.com/hajimehoshi/ebiten/v2"
)

type Mob struct {
	Name string
	components.PositionComponent
	components.HealthComponent
	components.MagicComponent
	Blocked bool

	image *ebiten.Image
}

func NewMob(n string, x, y int, i *ebiten.Image, blocked bool) *Mob {
	m := &Mob{
		Name:    n,
		Blocked: blocked,
		image:   i,
	}
	m.X = x
	m.Y = y

	return m
}

func (m *Mob) GetPosition() (x, y int) {
	return m.X, m.Y
}

func (m *Mob) Move(x int, y int, g *grid.Grid) {
	if g.TileAt(x, y) == nil || !g.TileAt(x, y).Blocked {
		m.X = x
		m.Y = y
	}
	//log.Debug().Msgf("Move - MobX: %d - MobY: %d", m.X, m.Y)
}

func (m *Mob) Image() *ebiten.Image {
	return m.image
}
