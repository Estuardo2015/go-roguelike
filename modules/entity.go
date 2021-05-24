package modules

import (
	"github.com/Estuardo2015/rogue_wizard/modules/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	GetPosition() (x, y int)
	Move(x, y int, g *Grid)
	Image() *ebiten.Image
}

type Mob struct {
	Name string
	PositionComponent
	HealthComponent
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

func (m *Mob) Move(x int, y int, g *Grid) {
	if g.TileAt(x, y) == nil || !g.TileAt(x, y).Blocked {
		m.X = x
		m.Y = y
	}
	//log.Debug().Msgf("Move - MobX: %d - MobY: %d", m.X, m.Y)
}

func (m *Mob) Image() *ebiten.Image {
	return m.image
}

func MoveEntities(g *Game) {
	for _, e := range g.Level.Entities {
		x, y := e.GetPosition()
		switch pickDirection() {
		case 0: // Up
			e.Move(x, y-1, g.Level.Grid)
		case 1: // Down
			e.Move(x, y+1, g.Level.Grid)
		case 2: // Left
			e.Move(x-1, y, g.Level.Grid)
		case 3: // Right
			e.Move(x+1, y, g.Level.Grid)
		}
	}
}

func pickDirection() int {
	return utils.RandInt(4)
}
