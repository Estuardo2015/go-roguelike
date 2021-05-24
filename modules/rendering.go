package modules

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rs/zerolog/log"
)

func RenderGame(g *Game, screen *ebiten.Image) {
	g.Level.Camera.LookAt(g, screen)
}

func ScreenDraw(x, y int, img, screen *ebiten.Image, g *Game) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x*g.Level.TileWidth), float64(y*g.Level.TileWidth))
	screen.DrawImage(img, op)
}

func LogCursor() {
	x, y := ebiten.CursorPosition()
	log.Debug().Msgf("CursorX: %d CursorY: %d", x, y)
}
