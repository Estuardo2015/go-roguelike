package modules

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type Camera struct {
	Target Entity

	ScreenWidth  int
	ScreenHeight int
	LevelWidth   int
	LevelHeight  int
	TileWidth    int
}

func NewCamera(e Entity, sw, sh, lw, lh int) *Camera {
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

	for x := 0; x < c.ScreenWidth; x++ {
		for y := 0; y < c.ScreenHeight; y++ {
			tile := GetTile(g.Level, x+camX, y+camY)
			if tile == nil {
				continue
			}

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*g.Level.TileWidth), float64(y*g.Level.TileWidth))
			screen.DrawImage(tile.Image, op)
		}
	}

	for _, e := range g.Level.Entities {
		eX, eY := e.GetPosition()
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64((eX-camX)*g.Level.TileWidth), float64((eY-camY)*g.Level.TileWidth))
		screen.DrawImage(e.Image(), op)
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64((g.Level.Player.X-camX)*g.Level.TileWidth), float64((g.Level.Player.Y-camY)*g.Level.TileWidth))
	screen.DrawImage(g.Level.Player.Image(), op)
}
