package modules

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func HandleInput(g *Game) {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Level.Player.Move(g.Level.Player.X, g.Level.Player.Y-1, g.Level.TileGrid)
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Level.Player.Move(g.Level.Player.X, g.Level.Player.Y+1, g.Level.TileGrid)
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Level.Player.Move(g.Level.Player.X-1, g.Level.Player.Y, g.Level.TileGrid)
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Level.Player.Move(g.Level.Player.X+1, g.Level.Player.Y, g.Level.TileGrid)
	}
}
