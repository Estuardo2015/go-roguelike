package modules

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func RenderGame(g *Game, screen *ebiten.Image) {
	g.Level.Camera.LookAt(g, screen)
}
