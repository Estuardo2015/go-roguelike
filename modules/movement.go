package modules

func Move(g *Game, e *Entity, x, y int) {
	if g.Level.TileMap[x][y] == nil || !g.Level.TileMap[x][y].Blocked {
		ClearEntity(g, e.X, e.Y)
		g.Level.EntityMap[x][y] = e
		e.X = x
		e.Y = y
	}
}
