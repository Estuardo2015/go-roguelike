package modules

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

var WallImg, AirImg, PlayerImg *ebiten.Image

// Load sprites
func init() {
	var err error

	PlayerImg, _, err = ebitenutil.NewImageFromFile("./assets/player.png")
	if err != nil {
		log.Fatal(err)
	}

	WallImg, _, err = ebitenutil.NewImageFromFile("./assets/wall.png")
	if err != nil {
		log.Fatal(err)
	}

	AirImg, _, err = ebitenutil.NewImageFromFile("./assets/air.png")
	if err != nil {
		log.Fatal(err)
	}
}
