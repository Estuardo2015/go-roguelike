package modules

import "github.com/hajimehoshi/ebiten/v2"

func RenderGame(g *Game, screen *ebiten.Image) {
	DrawEnvironment(g, screen)
	DrawPlayer(g, screen)
	DrawEntities(g, screen)
}

func DrawEnvironment(g *Game, screen *ebiten.Image) {
	for x := 0; x < g.Level.Width; x++ {
		for y := 0; y < g.Level.Height; y++ {
			if g.Level.TileGrid[x][y] != nil {
				tile := g.Level.TileGrid[x][y]
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(tile.X*g.Level.TileWidth), float64(tile.Y*g.Level.TileHeight))
				screen.DrawImage(tile.Image, op)
			}
		}
	}
}

func DrawPlayer(g *Game, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.Level.Player.X*g.Level.TileWidth), float64(g.Level.Player.Y*g.Level.TileHeight))
	screen.DrawImage(g.Level.Player.Image, op)
}

func DrawEntities(g *Game, screen *ebiten.Image) {
	for _, e := range g.Level.Entities {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(e.X*g.Level.TileWidth), float64(e.Y*g.Level.TileHeight))
		screen.DrawImage(g.Level.Player.Image, op)
	}
}
