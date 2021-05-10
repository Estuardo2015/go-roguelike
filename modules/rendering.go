package modules

import "github.com/hajimehoshi/ebiten/v2"

func DrawLevel(g *Game, screen *ebiten.Image) {
	for x := 0; x < g.Level.Width; x++ {
		for y := 0; y < g.Level.Height; y++ {
			// Draw tiles
			if g.Level.TileMap[x][y] != nil {
				tile := g.Level.TileMap[x][y]
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(tile.X*g.Level.TileWidth), float64(tile.Y*g.Level.TileHeight))
				screen.DrawImage(tile.Image, op)
			} else if g.Level.EntityMap[x][y] != nil {
				// Draw Entities
				entity := g.Level.EntityMap[x][y]
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(entity.X*g.Level.TileWidth), float64(entity.Y*g.Level.TileHeight))
				screen.DrawImage(entity.Image, op)
			}
		}
	}
}

func ClearEntity(g *Game, x, y int) {
	g.Level.EntityMap[x][y] = nil
}
