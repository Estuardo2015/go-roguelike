package modules

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rs/zerolog/log"
	"math/rand"
	"time"
)

type Entity interface {
	GetPosition() (x, y int)
	Move(x, y int, tg [][]*Tile)
	Image() *ebiten.Image
}

type Mob struct {
	Name    string
	X       int
	Y       int
	Blocked bool

	image *ebiten.Image
}

func (m *Mob) GetPosition() (x, y int) {
	return m.X, m.Y
}

func (m *Mob) Move(x int, y int, tg [][]*Tile) {
	if tg[x][y] == nil || !tg[x][y].Blocked {
		m.X = x
		m.Y = y
	}
	log.Debug().Msgf("Move - MobX: %d - MobY: %d", m.X, m.Y)
}

func (m *Mob) Image() *ebiten.Image {
	return m.image
}

func MoveEntities(g *Game) {
	for _, e := range g.Level.Entities {
		if willMove() {
			x, y := e.GetPosition()
			switch pickDirection() {
			case 0: // Up
				e.Move(x, y-1, g.Level.TileGrid)
			case 1: // Down
				e.Move(x, y+1, g.Level.TileGrid)
			case 2: // Left
				e.Move(x-1, y, g.Level.TileGrid)
			case 3: // Right
				e.Move(x+1, y, g.Level.TileGrid)
			}
		}
	}
}

func willMove() bool {
	// Flip a coin to decide if entity will move
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	willMoveInt := r.Intn(100)
	log.Debug().Msgf("willMove - %d", willMoveInt)
	if willMoveInt > 95 {
		return true
	} else {
		return false
	}
}

func pickDirection() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	dir := r.Intn(4)
	return dir
}
