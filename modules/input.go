package modules

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func HandleInput(g *Game) {
	// Write your game's logical update.
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		Move(g, g.Level.Player, g.Level.Player.X, g.Level.Player.Y-1)
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		Move(g, g.Level.Player, g.Level.Player.X, g.Level.Player.Y+1)
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		Move(g, g.Level.Player, g.Level.Player.X-1, g.Level.Player.Y)
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		Move(g, g.Level.Player, g.Level.Player.X+1, g.Level.Player.Y)
	}
}
