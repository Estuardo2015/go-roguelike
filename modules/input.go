package modules

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rs/zerolog/log"
)

func HandleInput(g *Game) {
	HandleKeys(g)
	HandleMouse(g)
}

func HandleKeys(g *Game) {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Level.Player.Move(g.Level.Player.X, g.Level.Player.Y-1, g.Level.Grid)
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Level.Player.Move(g.Level.Player.X, g.Level.Player.Y+1, g.Level.Grid)
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Level.Player.Move(g.Level.Player.X-1, g.Level.Player.Y, g.Level.Grid)
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Level.Player.Move(g.Level.Player.X+1, g.Level.Player.Y, g.Level.Grid)
	}
}

func HandleMouse(g *Game) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := GetCursorXY()
		e := g.Level.EntityAt(x, y)

		if e != nil {
			eX, eY := e.GetPosition()
			log.Debug().Msgf("Entity clicked at (%d,%d)", eX, eY)
		}
	}
}
