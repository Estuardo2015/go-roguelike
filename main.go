package main

import (
	"github.com/Estuardo2015/rogue_wizard/modules"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rs/zerolog/log"
)

const (
	WindowWidth  = 840
	WindowHeight = 400
	MaxTPS       = 10
)

func main() {
	game := modules.NewGame()

	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowResizable(false)
	ebiten.SetWindowTitle("Rogue Wizard")
	ebiten.SetMaxTPS(MaxTPS)
	// Start game loop
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal().Msg(err.Error())
	}
}
