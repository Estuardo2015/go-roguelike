package modules

import (
	"github.com/Estuardo2015/rogue_wizard/modules/commons"
	_ "image/png"
)

type Level struct {
	*Grid

	Player   *Player
	Entities []Entity

	Camera *Camera
}

func (l *Level) AddEntity(e Entity) {
	l.Entities = append(l.Entities, e)
}

func NewLevel(tw int, g *Grid, p *Player) *Level {
	l := &Level{
		Grid:   g,
		Player: p,
	}

	// Initialize camera
	l.Camera = NewCamera(l.Player, commons.ScreenWidthTiles, commons.ScreenHeightTiles, l.Width(), l.Height())

	return l
}
