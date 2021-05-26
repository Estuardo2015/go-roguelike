package modules

import (
	"github.com/Estuardo2015/rogue_wizard/modules/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type Camera struct {
	Target entity.Entity

	ScreenWidth  int
	ScreenHeight int
	LevelWidth   int
	LevelHeight  int
	TileWidth    int
}

func NewCamera(e entity.Entity, sw, sh, lw, lh int) *Camera {
	return &Camera{
		Target:       e,
		ScreenWidth:  sw,
		ScreenHeight: sh,
		LevelWidth:   lw,
		LevelHeight:  lh,
	}
}

func (c Camera) CalculateOrigin() (X, Y int) {
	tX, tY := c.Target.GetPosition()

	camX := int(math.Max(0, math.Min(float64(tX-(c.ScreenWidth/2)), float64(c.LevelWidth-c.ScreenWidth))))
	camY := int(math.Max(0, math.Min(float64(tY-(c.ScreenHeight/2)), float64(c.LevelHeight-c.ScreenHeight))))

	return camX, camY
}

func (c Camera) LookAt(g *Game, screen *ebiten.Image) {
	camX, camY := c.CalculateOrigin()

	// Draw world
	for tgX := 0; tgX < c.ScreenWidth; tgX++ {
		for tgY := 0; tgY < c.ScreenHeight; tgY++ {
			tile := g.Level.TileAt(tgX+camX, tgY+camY)
			if tile == nil {
				continue
			}

			ScreenDraw(tgX, tgY, tile.Image, screen, g)
		}
	}

	// Draw Entities
	g.Level.ForEachEntity(func(e entity.Entity) {
		eX, eY := e.GetPosition()
		ScreenDraw(eX-camX, eY-camY, e.Image(), screen, g)
	})

	// Draw Player
	ScreenDraw(g.Level.Player.X-camX, g.Level.Player.Y-camY, g.Level.Player.Image(), screen, g)
}
