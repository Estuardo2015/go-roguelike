package modules

import "github.com/hajimehoshi/ebiten/v2"

type Entity struct {
	Name    string
	X       int
	Y       int
	Image   *ebiten.Image
	Blocked bool
}
