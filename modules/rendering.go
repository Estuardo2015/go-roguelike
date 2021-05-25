package modules

import (
	"github.com/Estuardo2015/rogue_wizard/modules/commons"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rs/zerolog/log"
	"math"
)

func RenderGame(g *Game, screen *ebiten.Image) {
	g.Level.Camera.LookAt(g, screen)
}

func ScreenDraw(x, y int, img, screen *ebiten.Image, g *Game) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x*commons.TileWidth), float64(y*commons.TileWidth))
	screen.DrawImage(img, op)
}

func GetCursorXY() (tileX, tileY int) {
	x, y := ebiten.CursorPosition()
	if (x >= 0) && (y >= 0) {
		tileXF64 := math.Floor(float64(x) / float64(commons.TileWidth))
		tileYF64 := math.Floor(float64(y) / float64(commons.TileWidth))

		tileX := int(tileXF64)
		tileY := int(tileYF64)

		log.Debug().Msgf("Cursor: (%d,%d)", tileX, tileY)
		return tileX, tileY
	}
	return -1, -1
}
