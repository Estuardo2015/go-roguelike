package modules

func Move(g *Game, e *Entity, targetX, targetY int) {
	if CanMoveTo(g, targetX, targetY) {
		e.X = targetX
		e.Y = targetY
	}
}

func CanMoveTo(g *Game, targetX int, targetY int) bool {
	return g.Level.TileGrid[targetX][targetY] == nil ||
		!g.Level.TileGrid[targetX][targetY].Blocked
}
