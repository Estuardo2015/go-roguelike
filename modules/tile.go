package modules

import "github.com/hajimehoshi/ebiten/v2"

type Tile struct {
	X       int
	Y       int
	Image   *ebiten.Image
	Blocked bool
}
