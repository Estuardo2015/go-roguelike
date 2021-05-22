package main

import (
	"github.com/Estuardo2015/rogue_wizard/modules"
	"github.com/Estuardo2015/rogue_wizard/modules/commons"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rs/zerolog/log"
)

func main() {
	game := modules.NewGame()

	ebiten.SetWindowSize(commons.WindowWidth, commons.WindowHeight)
	ebiten.SetWindowResizable(false)
	ebiten.SetWindowTitle(commons.WindowTitle)
	ebiten.SetMaxTPS(commons.MaxTPS)
	// Start game loop
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal().Msg(err.Error())
	}
}
