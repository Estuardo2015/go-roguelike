package grid

type Grid struct {
	width    int // tiles wide
	height   int // tiles high
	tileGrid [][]*Tile
}

func (g Grid) TileAt(x, y int) *Tile {
	if x < 0 || x >= g.width || y < 0 || y >= g.height {
		return nil
	}

	return g.tileGrid[x][y]
}

func (g Grid) Width() int {
	return g.width
}

func (g Grid) Height() int {
	return g.height
}

func NewGrid(w, h int, tg [][]*Tile) *Grid {
	g := &Grid{
		width:    len(tg),
		height:   len(tg[0]),
		tileGrid: tg,
	}

	return g
}
