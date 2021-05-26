package modules

import (
	"github.com/Estuardo2015/rogue_wizard/modules/commons"
	"github.com/Estuardo2015/rogue_wizard/modules/entity"
	"github.com/Estuardo2015/rogue_wizard/modules/grid"
	_ "image/png"
)

type Level struct {
	*grid.Grid

	Player *Player
	entity.Manager

	Camera *Camera
}

func NewLevel(tw int, g *grid.Grid, p *Player) *Level {
	l := &Level{
		Grid:   g,
		Player: p,
	}

	// Initialize camera
	l.Camera = NewCamera(l.Player, commons.ScreenWidthTiles, commons.ScreenHeightTiles, l.Width(), l.Height())

	return l
}
