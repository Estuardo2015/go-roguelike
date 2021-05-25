package modules

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

var WallImg,
	PlayerImg,
	GrassImg,
	ZombieImg *ebiten.Image

// Load sprites
func init() {
	var err error

	PlayerImg, _, err = ebitenutil.NewImageFromFile("./assets/img/player.png")
	if err != nil {
		log.Fatal(err)
	}

	WallImg, _, err = ebitenutil.NewImageFromFile("./assets/img/wall.png")
	if err != nil {
		log.Fatal(err)
	}

	GrassImg, _, err = ebitenutil.NewImageFromFile("./assets/img/grass.png")
	if err != nil {
		log.Fatal(err)
	}

	ZombieImg, _, err = ebitenutil.NewImageFromFile("./assets/img/zombie.png")
	if err != nil {
		log.Fatal(err)
	}
}
