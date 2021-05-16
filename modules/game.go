package modules

import (
	"github.com/Estuardo2015/rogue_wizard/commons"
	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements ebiten.Game interface.
type Game struct {
	Level *Level
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	HandleInput(g)
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	RenderGame(g, screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return commons.ScreenWidthPixels, commons.ScreenHeightPixels
}

//NewGame creates a new Game Object and initializes the data
//This is a pretty solid refactor candidate for later
func NewGame() *Game {
	g := &Game{}
	p := NewPlayer("Hiro", 20, 20, PlayerImg)

	g.Level = NewLevel(80, 50, commons.TileWidth, p)

	BuildDemoLevel(g.Level)

	return g
}
