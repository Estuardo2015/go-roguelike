package entity

import (
	"github.com/Estuardo2015/rogue_wizard/modules/grid"
	"github.com/Estuardo2015/rogue_wizard/modules/utils"
)

type Manager struct {
	entities []Entity
}

func (em *Manager) AddEntity(e Entity) {
	em.entities = append(em.entities, e)
}

func (em *Manager) RemoveEntity() {

}

func (em *Manager) EntityAt(x, y int) Entity {
	for _, e := range em.entities {
		eX, eY := e.GetPosition()
		if (eX == x) && (eY == y) {
			return e
		}
	}

	return nil
}

func (em *Manager) ForEachEntity(entityFunc func(e Entity)) {
	for _, e := range em.entities {
		entityFunc(e)
	}
}

func (em *Manager) MoveEntities(g *grid.Grid) {
	for _, e := range em.entities {
		x, y := e.GetPosition()
		switch pickDirection() {
		case 0: // Up
			e.Move(x, y-1, g)
		case 1: // Down
			e.Move(x, y+1, g)
		case 2: // Left
			e.Move(x-1, y, g)
		case 3: // Right
			e.Move(x+1, y, g)
		}
	}
}

func pickDirection() int {
	return utils.RandInt(4)
}
