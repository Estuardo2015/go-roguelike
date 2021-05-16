package modules

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func RenderGame(g *Game, screen *ebiten.Image) {
	g.Level.Camera.LookAt(g, screen)
}

func DrawTile(g *Game, screen *ebiten.Image, tile *Tile, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x*g.Level.TileWidth), float64(y*g.Level.TileWidth))
	screen.DrawImage(tile.Image, op)
}

func DrawPlayer(g *Game, screen *ebiten.Image, p *Player, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x*g.Level.TileWidth), float64(x*g.Level.TileWidth))
	screen.DrawImage(g.Level.Player.Image, op)
}
