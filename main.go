package main

import (
	"github.com/Estuardo2015/rogue_wizard/modules"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rs/zerolog/log"
)

func main() {
	game := modules.NewGame()

	ebiten.SetWindowResizable(false)
	ebiten.SetWindowTitle("Rogue Wizard")
	ebiten.SetMaxTPS(10)
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal().Msg(err.Error())
	}
}
